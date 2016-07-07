package gocensys

//QueryJob represents an SQL query job.
type QueryJob struct {
	Status        string `json:"status"`
	Configuration struct {
		Query string `json:"query"`
	} `json:"configuration"`
	JobID string `json:"job_id"`
}

//QueryJobStatus represents the status of an SQL query job that has already been started.
type QueryJobStatus struct {
	Status        string `json:"status"`
	Configuration struct {
		Query string `json:"query"`
	} `json:"configuration"`
	JobID string `json:"job_id"`
	Error string `json:"error"`
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
func (c CensysAPI) GetQueryJobStatus(jobID string) (QueryJobStatus, error) {
	var job QueryJobStatus

	err := c.apiGet("/query/"+jobID, nil, &job)

	return job, err
}
