package gocensys

import (
	"testing"
	"time"
)

func Test_Export(t *testing.T) {
	data := map[string]interface{}{
		"query":   "SELECT location.country, count(ip) FROM ipv4.20151020 GROUP BY location.country;",
		"format":  "json",
		"flatten": false,
	}

	job, err := api.StartExportJob(data)
	if err != nil {
		t.Fatal(err)
	}

	//Jobs typically require 15-30 seconds to execute.
	time.Sleep(30 * time.Second)

	_, err = api.GetExportJobStatus(job.JobID)
	if err != nil {
		t.Fatal(err)
	}
}
