package gocensys

import "testing"

func Test_Search(t *testing.T) {
	query := map[string]interface{}{
		"query":  "*",
		"page":   2,
		"fields": []string{"ip", "location.country", "autonomous_system.asn"},
	}

	s, err := api.Search("ipv4", query)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(s)
	}
}
