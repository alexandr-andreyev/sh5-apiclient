package sh5apiclient

type Sh5BaseRequest struct {
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

type sh5BaseResponse struct {
	UserName   string `json:"UserName"`
	ErrorCode  int    `json:"errorCode"`
	ErrMessage string `json:"errMessage"`
	Version    string `json:"Version"`
	Actioname  string `json:"actionName"`
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

// proc request struct
type Sh5ProcRequest struct {
	Sh5BaseRequest
	ProcName string
	Input    []ShInputData
}

// inputdata for proc
type ShInputData struct {
	Head     string   `json:"head"`
	Original []string `json:"original"`
	Values   [][]any  `json:"values"`
}

// sh5 proc response
type Sh5ProcResponse struct {
	sh5BaseResponse
	ShTable []ShTable `json:"shTable"`
}

type ShTable struct {
	Head     string   `json:"head"`
	RecCount int      `json:"recCount"`
	Original []string `json:"original"`
	Fields   []string `json:"fields"`
	Values   [][]any  `json:"values"`
}

type Sh5ProcParseResponse struct {
	sh5BaseResponse
	Data map[string][]map[string]string
}
