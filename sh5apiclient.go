package sh5apiclient

import (
	"bytes"
	"encoding/json"
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

func (c *Client) newRequest(method, url string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (с *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := с.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
