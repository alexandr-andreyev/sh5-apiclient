package sh5apiclient

import (
	"fmt"
)

// Request server settings and database information
func (c Client) Sh5Info() (*sh5SHInfoResponse, error) {
	url := fmt.Sprintf("%s:%d/api/sh5info", c.BaseURL, c.Port)

	sh5data := Sh5BaseRequest{
		Username: c.UserName,
		Password: c.Password,
	}

	req, err := c.newRequest("POST", url, sh5data)
	if err != nil {
		return nil, err
	}

	var shInfo sh5SHInfoResponse
	_, err = c.do(req, &shInfo)
	return &shInfo, err
}
