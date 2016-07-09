package gocensys

import (
	"testing"
	"time"
)

func Test_QueryJob(t *testing.T) {
	queryJob, err := api.StartQueryJob("select p110.pop3.ssl_2.certificate.parsed.issuer.province from ipv4.20160708 limit 1000;")
	if err != nil {
		t.Fatal(err)
	}

	//Jobs typically require 15-30 seconds to execute.
	time.Sleep(30 * time.Second)

	_, err = api.GetQueryJobStatus(queryJob.JobID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.GetQueryJobResults(queryJob.JobID, "1")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_QuerySeries(t *testing.T) {
	_, err := api.GetQuerySeries()
	if err != nil {
		t.Fatal(err)
	}

	_, err = api.GetQuerySeriesDetails("ipv4")
	if err != nil {
		t.Fatal(err)
	}
}
