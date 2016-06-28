package gocensys

import "testing"

func Test_ViewDocument(t *testing.T) {
	_, err := api.ViewDocument("websites", "google.com")
	if err != nil {
		t.Fatal(err)
	}
}
