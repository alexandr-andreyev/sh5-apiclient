package sh5apiclient

import (
	"errors"
	"fmt"
)

func (c Client) Sh5Exec(procName string, inputData []ShInputData) (*Sh5ProcResponse, error) {
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

func (c Client) Sh5ExecWithParse(req *Sh5ProcResponse, fields bool) (*Sh5ProcParseResponse, error) {
	resultResponse := Sh5ProcParseResponse{}
	resultResponse.sh5BaseResponse = req.sh5BaseResponse

	if req.ErrorCode == 0 {
		result := make(map[string][]map[string]string)

		for _, headDataset := range req.ShTable {
			if len(headDataset.Values) == 0 {
				continue
			}
			valuesArr := make([]map[string]string, headDataset.RecCount)
			for keyValuesMap, _ := range headDataset.Values[0] {
				valuesArr[keyValuesMap] = make(map[string]string)
			}

			for keyField, fieldName := range headDataset.Fields {
				for valueId, value := range headDataset.Values[keyField] {
					valuesArr[valueId][fieldName] = fmt.Sprintf("%v", value)
				}
			}
			head := fmt.Sprintf("Head%s", headDataset.Head)
			result[head] = append(result[headDataset.Head], valuesArr...)
		}

		resultResponse.Data = result
		return &resultResponse, nil
	}

	return &resultResponse, errors.New(req.ErrMessage)
}
