package sh5apiclient

import "fmt"

// Request server settings and database information
type ODocsInput struct {
	FromDate string
	ToDate   string
	Sunit    string
	Corr     string
}

func (c Client) ReadODocs(input ODocsInput) (*Sh5ProcParseResponse, error) {
	var inputData []ShInputData

	// Фильтр по датам - обязательный, всегда добавляем
	if input.FromDate != "" || input.ToDate != "" {
		filterDate := ShInputData{
			Head:     "108",
			Original: []string{"1", "2", "30", "225#2\\1"},
			Values:   [][]interface{}{{input.FromDate}, {input.ToDate}, {"480"}, {}}, // 480 ?
		}
		inputData = append(inputData, filterDate)
	}

	// Фильтр по месту реализации - добавляем только если указан
	if input.Sunit != "" {
		filterSUnit := ShInputData{
			Head:     "226#10",
			Original: []string{"1"},
			Values:   [][]interface{}{{input.Sunit}},
		}
		inputData = append(inputData, filterSUnit)
	}

	// Фильтр по корреспонденту - добавляем только если указан
	if input.Corr != "" {
		filterCorr := ShInputData{
			Head:     "107#10",
			Original: []string{"1"},
			Values:   [][]interface{}{{input.Corr}},
		}
		inputData = append(inputData, filterCorr)
	}

	// Если нет ни одного фильтра, возвращаем ошибку или создаем минимальный запрос
	if len(inputData) == 0 {
		return nil, fmt.Errorf("at least one filter parameter must be provided")
	}

	resp, err := c.Sh5ExecPrettyJson("ODocs", inputData, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
