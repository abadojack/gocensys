package gocensys

import "testing"

func Test_CreateReport(t *testing.T) {
	index := "websites"
	query := map[string]interface{}{
		"query":   "80.http.get.headers.server: nginx",
		"field":   "location.country",
		"buckets": 100,
	}

	_, err := api.CreateReport(index, query)
	if err != nil {
		t.Fatal(err)
	}
}
