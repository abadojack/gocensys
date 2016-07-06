package gocensys

//Job represents an export job.
type Job struct {
	Status        string `json:"status"`
	Configuration struct {
		Format    string `json:"format"`
		Compress  bool   `json:"compress"`
		Headers   bool   `json:"headers"`
		Delimiter bool   `json:"delimiter"`
		Flatten   bool   `json:"flatten"`
		Query     string `json:"query"`
	} `json:"configuration"`
	JobID string `json:"job_id"`
}

//JobStatus represents the status of a submitted job.
type JobStatus struct {
	Status     string `json:"status"`
	Statistics struct {
		BackendTime interface{} `json:"backend_time"`
		ResultCount interface{} `json:"result_count"`
	} `json:"statistics"`
	JobID         string `json:"job_id"`
	Configuration struct {
		Format    string `json:"format"`
		Compress  bool   `json:"compress"`
		Headers   bool   `json:"headers"`
		Delimiter bool   `json:"delimiter"`
		Flatten   bool   `json:"flatten"`
		Query     string `json:"query"`
	} `json:"configuration"`
	DownloadPaths []string `json:"download_paths"`
}

//StartExportJob lets you submit an export job.
func (c CensysAPI) StartExportJob(data map[string]interface{}) (Job, error) {
	var job Job

	err := c.apiPost("/export", data, &job)

	return job, err
}

//GetJobStatus lets you retrieve information about a previously submitted job.
func (c CensysAPI) GetJobStatus(jobID string) (JobStatus, error) {
	var jobStatus JobStatus

	err := c.apiGet("/export/"+jobID, nil, &jobStatus)

	return jobStatus, err
}
