package sh5apiclient

type Sh5BaseRequest struct {
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

type sh5BaseResponse struct {
	Sh5BaseRequest
	ErrorCode  int    `json:"errorCode"`
	ErrMessage string `json:"errMessage"`
	Version    string `json:"Version"`
	//Actioname  string `json:"actionName"`
	//ActionType string `json:"actionType"`
}

type sh5SHInfoResponse struct {
	sh5BaseResponse
	LinkDisp string `json:"LinkDisp"`
	TimeOut  int    `json:"timeout"`
	SH5DB    `json:"DB"`
}

type SH5DB struct {
	Ident   string `json:"Ident"`
	Size    string `json:"Size"`
	Version string `json:"Version"`
}
