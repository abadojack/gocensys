package gocensys

// GetSeries endpoint returns a data on the types of scans censys regularly performs ("series").
func (c CensysAPI) GetSeries() (RawData, error) {
	var rawData RawData
	err := c.apiGet("/data", nil, &rawData)

	return rawData, err
}

// ViewSeries endpoint returns data censys has about a particular series —
// a scan of the same protocol and destination across time—including the list of scans.
func (c CensysAPI) ViewSeries(series string) (RawSeriesDetail, error) {
	var rawSeriesDetail RawSeriesDetail
	err := c.apiGet("/data/"+series, nil, &rawSeriesDetail)

	return rawSeriesDetail, err
}

// ViewResult endpoint returns data on a particular scan ("result"), as found in the Get Series or View Series endpoints.
func (c CensysAPI) ViewResult(series, result string) (ScanResult, error) {
	var scanResult ScanResult
	err := c.apiGet("/data/"+series+"/"+result, nil, &scanResult)

	return scanResult, err
}

// RawData represents raw data that can be downloaded from Censys.
type RawData struct {
	PrimarySeries PrimarySeries `json:"primary_series"`
	RawSeries     RawSeries     `json:"raw_series"`
}
