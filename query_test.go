package gocensys

import (
	"testing"
	"time"
)

func Test_Query(t *testing.T) {
	queryJob, err := api.StartQueryJob("select ip from ipv4.20150902 limit 1000;")
	if err != nil {
		t.Fatal(err)
	}

	//Jobs typically require 15-30 seconds to execute.
	time.Sleep(30 * time.Second)

	_, err = api.GetQueryJobStatus(queryJob.JobID)
	if err != nil {
		t.Fatal(err)
	}
}
