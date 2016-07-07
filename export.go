package gocensys

//ExportJob represents an export job.
type ExportJob struct {
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

//ExportJobStatus represents the status of a submitted job.
type ExportJobStatus struct {
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
func (c CensysAPI) StartExportJob(data map[string]interface{}) (ExportJob, error) {
	var job ExportJob

	err := c.apiPost("/export", data, &job)

	return job, err
}

//GetExportJobStatus lets you retrieve information about a previously submitted job.
func (c CensysAPI) GetExportJobStatus(jobID string) (ExportJobStatus, error) {
	var jobStatus ExportJobStatus

	err := c.apiGet("/export/"+jobID, nil, &jobStatus)

	return jobStatus, err
}
