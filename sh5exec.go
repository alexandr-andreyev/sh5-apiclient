package sh5apiclient

import (
	"errors"
	"fmt"
)

func (c Client) Sh5ExecRaw(procName string, inputData []ShInputData) (*Sh5ProcResponse, error) {
	input := Sh5ProcRequest{
		Sh5BaseRequest: Sh5BaseRequest{
			Username: c.UserName,
			Password: c.Password,
		},
		ProcName: procName,
		Input:    inputData,
	}

	req, err := c.newRequest("POST", c.BaseURL, input)
	if err != nil {
		return nil, err
	}

	var sh5Response Sh5ProcResponse
	_, err = c.do(req, &sh5Response)

	return &sh5Response, err
}

func (c Client) Sh5ExecPrettyJson(procName string, inputData []ShInputData, fields bool) (*Sh5ProcParseResponse, error) {
	req, err := c.Sh5ExecRaw(procName, inputData)
	if err != nil {
		return nil, err
	}

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
					switch value.(type) {
					case float64:
						// Привести к int
						newValue := int(value.(float64))
						valuesArr[valueId][fieldName] = fmt.Sprintf("%v", newValue)
					default:
						valuesArr[valueId][fieldName] = fmt.Sprintf("%v", value)
					}

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
