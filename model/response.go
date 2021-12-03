package model

type RequestResponse struct {
	Data    interface{} `form:"data" json:"data" xml:"data"`
	Count   int         `form:"count" json:"count" xml:"count"`
	Error   string      `form:"error" json:"error" xml:"error"`
	Message string      `form:"message" json:"message" xml:"message"`
	Result  bool        `form:"result" json:"result" xml:"result"`
}

type Response struct {
	Message interface{} `form:"message" json:"message" xml:"message"`
	Status  int         `form:"status" json:"status" xml:"status"`
	Error   string      `form:"error" json:"error" xml:"error"`
}
