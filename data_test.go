package gocensys

import "testing"

func Test_GetSeries(t *testing.T) {
	_, err := api.GetSeries()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ViewSeries(t *testing.T) {
	_, err := api.ViewSeries("22-ssh-banner-full_ipv4")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ViewResult(t *testing.T) {
	_, err := api.ViewResult("22-ssh-banner-full_ipv4", "20150930T0056")
	if err != nil {
		t.Fatal(err)
	}
}
