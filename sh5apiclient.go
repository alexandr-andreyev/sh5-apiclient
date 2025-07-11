package sh5apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	UserName   string
	Password   string
	HTTPClient *http.Client
}

func NewClient(baseURL string, port int, userName, password string, isSSL bool) *Client {
	scheme := "http"
	if isSSL {
		scheme = "https"
	}

	fullBaseURL := fmt.Sprintf("%s://%s:%d", scheme, baseURL, port)
	return &Client{
		BaseURL:  fullBaseURL,
		UserName: userName,
		Password: password,
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
