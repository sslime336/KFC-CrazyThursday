package pos

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

func End() *KFCPosResponse {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, getAMapQueryUrl(), nil)
	{
		req.Header.Set("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Host", "restapi.amap.com")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Connection", "keep-alive")
	}
	resp, err := client.Do(req)
	_ = err
	defer resp.Body.Close()
	r, _ := gzip.NewReader(resp.Body)
	data := make([]byte, 4096)
	n, _ := r.Read(data)
	data = data[:n]

	kpos := new(KFCPosResponse)
	_ = json.Unmarshal(data, kpos)
	return kpos
}
