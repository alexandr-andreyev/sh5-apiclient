package sh5apiclient

import "fmt"

func (c Client) Sh5Exec(procName string, inputData []ShInputData) (any, error) {
	url := fmt.Sprintf("%s:%d/api/sh5exec", c.BaseURL, c.Port)

	input := Sh5ProcRequest{
		Sh5BaseRequest: Sh5BaseRequest{
			Username: c.UserName,
			Password: c.Password,
		},
		ProcName: procName,
		Input:    inputData,
	}

	req, err := c.newRequest("POST", url, input)
	if err != nil {
		return nil, err
	}

	var sh5Response Sh5ProcResponse
	_, err = c.do(req, &sh5Response)
	return &sh5Response, err
}
