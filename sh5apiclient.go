package sh5apiclient

import (
	"fmt"
	"net/http"
	"time"
)

type Sh5Client struct {
	BaseURL    string
	Port       int
	UserName   string
	Password   string
	HTTPClient *http.Client
}

// type errorResponse struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

// type successResponse struct {
// 	Code int         `json:"code"`
// 	Data interface{} `json:"data"`
// }

func NewClient(BaseURL string, Port int, UserName, Password string) *Sh5Client {
	return &Sh5Client{
		BaseURL:  BaseURL,
		UserName: UserName,
		Password: Password,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// Request server settings and database information
func (c Sh5Client) Sh5Info() error {
	url := fmt.Sprintf("%s:%d/api/sh5info", c.BaseURL, c.Port)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("response: %+v", resp.Body)
	return nil
}

func (c Sh5Client) Sh5Exec() {

}
