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
	Parent     *Childtype    `json:"parent"`    
	Address    string        `json:"address"`   
	Distance   string        `json:"distance"`  
	Pname      string        `json:"pname"`     
	Importance []interface{} `json:"importance"`
	BizEXT     []interface{} `json:"biz_ext"`   
	BizType    string        `json:"biz_type"`  
	Cityname   string        `json:"cityname"`  
	Type       string        `json:"type"`      
	Photos     []interface{} `json:"photos"`    
	Typecode   string        `json:"typecode"`  
	Shopinfo   string        `json:"shopinfo"`  
	Poiweight  []interface{} `json:"poiweight"` 
	Childtype  *Childtype    `json:"childtype"` 
	Adname     string        `json:"adname"`    
	Name       string        `json:"name"`      
	Location   string        `json:"location"`  
	Tel        string        `json:"tel"`       
	Shopid     []interface{} `json:"shopid"`    
	ID         string        `json:"id"`        
}

type Suggestion struct {
	Keywords []interface{} `json:"keywords"`
	Cities   []interface{} `json:"cities"`  
}

type Childtype struct {
	AnythingArray []interface{}
	String        *string
}
