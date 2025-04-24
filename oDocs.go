package sh5apiclient

// Request server settings and database information
type ODocsInput struct {
	FromDate string
	ToDate   string
	Sunit    string
	Corr     string
}

func (c Client) ReadODocs(input ODocsInput) (*Sh5ProcParseResponse, error) {
	// Filtering orders from the request by date
	filterDate := ShInputData{
		Head:     "108",
		Original: []string{"1", "2", "30", "225#2\\1"},
		Values:   [][]interface{}{{input.FromDate}, {input.ToDate}, {"480"}, {}}, // 480 ?
	}

	// Filtering orders by implementation location (Possible)
	filterSUnit := ShInputData{
		Head:     "226#10",
		Original: []string{"1"},
		Values:   [][]interface{}{{input.Sunit}},
	}

	// Filtering orders by correspondent
	filterCorr := ShInputData{
		Head:     "107#10",
		Original: []string{"1"},
		Values:   [][]interface{}{{input.Corr}},
	}

	inputData := []ShInputData{filterDate, filterSUnit, filterCorr}

	resp, err := c.Sh5ExecPrettyJson("ODocs", inputData, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
