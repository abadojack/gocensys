package gocensys

//QueryJob represents an SQL query job.
type QueryJob struct {
	Status        string `json:"status"`
	Configuration struct {
		Query string `json:"query"`
	} `json:"configuration"`
	JobID string `json:"job_id"`
	Error string `json:"error"`
}

// QueryJobResults represents results of a query after it has completed successfully.
type QueryJobResults struct {
	Statistics struct {
		Rows          float64 `json:"rows"`
		Pages         float64 `json:"pages"`
		DataProcessed float64 `json:"data_processed"`
	} `json:"statistics"`
	Configuration struct {
		Query  string `json:"query"`
		Page   int    `json:"page"`
		Schema struct {
			Fields []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Mode string `json:"mode"`
			} `json:"fields"`
		} `json:"schema"`
	} `json:"configuration"`
	Rows []interface{} `json:"rows"`
}

//QuerySeriesDetails represents details about a series, including the list of tables and schema for the series.
type QuerySeriesDetails struct {
	Definition struct {
		Fields []struct {
			Type   string `json:"type"`
			Name   string `json:"name"`
			Mode   string `json:"mode"`
			Fields []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Mode string `json:"mode"`
			} `json:"fields,omitempty"`
		} `json:"fields"`
	} `json:"definition"`
	Tables []string `json:"tables"`
}

//QuerySeries represents series that can queried through the SQL interface.
type QuerySeries struct {
	Series []string `json:"series"`
}

//StartQueryJob lets you to start a new asynchronous query job.
func (c CensysAPI) StartQueryJob(query string) (QueryJob, error) {
	queryMap := map[string]interface{}{
		"query": query,
	}

	var job QueryJob

	err := c.apiPost("/query", queryMap, &job)

	return job, err
}

//GetQueryJobStatus allows you to determine whether a job has completed.
func (c CensysAPI) GetQueryJobStatus(jobID string) (QueryJob, error) {
	var job QueryJob

	err := c.apiGet("/query/"+jobID, nil, &job)

	return job, err
}

//GetQueryJobResults allows you to retrieve results of a query after it has completed successfully.
func (c CensysAPI) GetQueryJobResults(jobID, page string) (QueryJobResults, error) {
	var results QueryJobResults

	err := c.apiGet("/query/"+jobID+"/"+page, nil, &results)

	return results, err
}

//GetQuerySeries returns a list of series that can queried through the SQL interface.
func (c CensysAPI) GetQuerySeries() (QuerySeries, error) {
	var querySeries QuerySeries

	err := c.apiGet("/query_definitions", nil, &querySeries)

	return querySeries, err
}

//GetQuerySeriesDetails returns details about a series, including the list of tables and schema for the series.
func (c CensysAPI) GetQuerySeriesDetails(series string) (QuerySeriesDetails, error) {
	var querySeriesDetails QuerySeriesDetails

	err := c.apiGet("/query_definitions/"+series, nil, &querySeriesDetails)

	return querySeriesDetails, err
}
