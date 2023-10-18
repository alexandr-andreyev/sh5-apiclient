package sh5apiclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	Port       int
	UserName   string
	Password   string
	HTTPClient *http.Client
}

type errorResponse struct {
	ErrorCode  int    `json:"errorCode"`
	ErrMessage string `json:"errMessage"`
	Version    string `json:"Version"`
	Username   string `json:"UserName"`
}

type successResponse struct {
	ErrorCode  int    `json:"errorCode"`
	ErrMessage string `json:"errMessage"`
	Version    string `json:"Version"`
	Username   string `json:"UserName"`
	Actioname  string `json:"actionName"`
	ActionType string `json:"actionType"`
}

func NewClient(BaseURL string, Port int, UserName, Password string) *Client {
	return &Client{
		BaseURL:  BaseURL,
		Port:     Port,
		UserName: UserName,
		Password: Password,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (с *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) //body
	if err != nil {
		return nil, err
	}

	var errResp errorResponse
	if err = json.NewDecoder(resp.Body).Decode(&errResp); err == nil {
		return nil, errors.New(errResp.ErrMessage)
	}
	if errResp.ErrorCode != 0 {
		return nil, errors.New(errResp.ErrMessage)
	}

	return body, nil
}
