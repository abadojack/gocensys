package gocensys

//Report represents a report on the breakdown of a field in the result set of a query.
type Report struct {
	Status  string `json:"status"`
	Results []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
	} `json:"results"`
	Metadata struct {
		Count            int    `json:"count"`
		BackendTime      int    `json:"backend_time"`
		NonnullCount     int    `json:"nonnull_count"`
		OtherResultCount int    `json:"other_result_count"`
		Buckets          int    `json:"buckets"`
		ErrorBound       int    `json:"error_bound"`
		Query            string `json:"query"`
	} `json:"metadata"`
}

//CreateReport lets you run aggregate reports on the breakdown of a field in a result set analogous to the "Build Report".
func (c CensysAPI) CreateReport(index string, query map[string]interface{}) (Report, error) {
	var report Report

	err := c.apiPost("/report/"+index, query, nil, &report)

	return report, err
}
