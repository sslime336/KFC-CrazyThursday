package pos

// PosResponse 为百度 IP 定位返回的结构
type PosResponse struct {
	Address string  `json:"address"`
	Content Content `json:"content"`
	Status  int64   `json:"status"`
}

type Content struct {
	AddressDetail AddressDetail `json:"address_detail"`
	Point         Point         `json:"point"`
	Address       string        `json:"address"`
}

type AddressDetail struct {
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	StreetNumber string `json:"street_number"`
	Adcode       string `json:"adcode"`
	Street       string `json:"street"`
	CityCode     int64  `json:"city_code"`
}

type Point struct {
	Y string `json:"y"`
	X string `json:"x"`
}

// KFCPosResponse 通过高德获取的附近肯德基位置
type KFCPosResponse struct {
	Suggestion Suggestion `json:"suggestion"`
	Count      string     `json:"count"`
	Infocode   string     `json:"infocode"`
	Pois       []Pois     `json:"pois"`
	Status     string     `json:"status"`
	Info       string     `json:"info"`
}

type Pois struct {
	Parent       *Childtype    `json:"parent"`
	Distance     string        `json:"distance"`
	Pcode        string        `json:"pcode"`
	Importance   []interface{} `json:"importance"`
	BizEXT       BizEXT        `json:"biz_ext"`
	Recommend    string        `json:"recommend"`
	Type         string        `json:"type"`
	Photos       []Photo       `json:"photos"`
	DiscountNum  string        `json:"discount_num"`
	Gridcode     string        `json:"gridcode"`
	Typecode     string        `json:"typecode"`
	Shopinfo     string        `json:"shopinfo"`
	Poiweight    []interface{} `json:"poiweight"`
	Citycode     string        `json:"citycode"`
	Adname       string        `json:"adname"`
	Children     []interface{} `json:"children"`
	Alias        []interface{} `json:"alias"`
	Tel          string        `json:"tel"`
	ID           string        `json:"id"`
	Tag          *Childtype    `json:"tag"`
	Event        []interface{} `json:"event"`
	EntrLocation *Childtype    `json:"entr_location"`
	IndoorMap    string        `json:"indoor_map"`
	Email        []interface{} `json:"email"`
	Timestamp    string        `json:"timestamp"`
	Website      []interface{} `json:"website"`
	Address      string        `json:"address"`
	Adcode       string        `json:"adcode"`
	Pname        string        `json:"pname"`
	BizType      string        `json:"biz_type"`
	Cityname     string        `json:"cityname"`
	Postcode     []interface{} `json:"postcode"`
	Match        string        `json:"match"`
	BusinessArea string        `json:"business_area"`
	IndoorData   IndoorData    `json:"indoor_data"`
	Childtype    *Childtype    `json:"childtype"`
	ExitLocation []interface{} `json:"exit_location"`
	Name         string        `json:"name"`
	Location     string        `json:"location"`
	Shopid       []interface{} `json:"shopid"`
	NaviPoiid    []interface{} `json:"navi_poiid"`
	GroupbuyNum  string        `json:"groupbuy_num"`
}

type BizEXT struct {
	Cost         string `json:"cost"`
	Rating       string `json:"rating"`
	MealOrdering string `json:"meal_ordering"`
}

type IndoorData struct {
	Cmsid     *Childtype `json:"cmsid"`
	Truefloor *Childtype `json:"truefloor"`
	Cpid      *Childtype `json:"cpid"`
	Floor     *Childtype `json:"floor"`
}

type Photo struct {
	Title []interface{} `json:"title"`
	URL   string        `json:"url"`
}

type Suggestion struct {
	Keywords []interface{} `json:"keywords"`
	Cities   []interface{} `json:"cities"`
}

type Childtype struct {
	AnythingArray []interface{}
	String        *string
}