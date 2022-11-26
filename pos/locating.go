package pos

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"

	_ "github.com/joho/godotenv/autoload"
)

func getBaiduQueryUrl() string {
	baiduKey := os.Getenv("BAIDUKEY")
	ip := getCurrentPubIP()
	return "https://api.map.baidu.com/location/ip?ak=" + baiduKey + "&ip=" + ip + "&coor=bd09ll"
}

func getAMapQueryUrl() string {
	aMapKey := os.Getenv("AMAPKEY")
	loc := getCurrentPos()
	return `https://restapi.amap.com/v3/place/around?key=` + aMapKey +
		`&location=` + loc + `&keywords=肯德基&types=050301&radius=1000&offset=20&page=1&extensions=base`
}

func getCurrentPos() string {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, getBaiduQueryUrl(), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:107.0) Gecko/20100101 Firefox/107.0")
	resp, err := client.Do(req)
	_ = err
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	baiduPos := new(PosResponse)
	_ = json.Unmarshal(data, baiduPos)
	return baiduPos.Content.Point.X + "," + baiduPos.Content.Point.Y
}

func getCurrentPubIP() string {
	rg := regexp.MustCompile(`(?:(?:1[0-9][0-9]\.)|(?:2[0-4][0-9]\.)|(?:25[0-5]\.)|(?:[1-9][0-9]\.)|(?:[0-9]\.)){3}(?:(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])|(?:[1-9][0-9])|(?:[0-9]))`)
	res := queryPubIP()
	return rg.FindString(res)
}

func queryPubIP() string {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://2022.ip138.com/", nil)
	{
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:107.0) Gecko/20100101 Firefox/107.0")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Referer", "https://www.ip138.com/")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("Sec-Fetch-Dest", "iframe")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("Pragma", "no-cache")
		req.Header.Set("Cache-Control", "no-cache")
	}
	resp, err := client.Do(req)
	_ = err
	defer resp.Body.Close()
	bodyText, _ := io.ReadAll(resp.Body)
	return string(bodyText)
}
