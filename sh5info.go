package sh5apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Sh5InfoRequest struct {
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

// Request server settings and database information
func (c Client) Sh5Info() ([]byte, error) {
	url := fmt.Sprintf("%s:%d/api/sh5info", c.BaseURL, c.Port)

	sh5data := Sh5InfoRequest{
		Username: c.UserName,
		Password: c.Password,
	}

	jsonSh5Data, _ := json.Marshal(&sh5data)
	paydata := bytes.NewReader(jsonSh5Data)

	req, err := http.NewRequest("POST", url, paydata)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
