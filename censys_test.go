package gocensys

import (
	"os"
	"testing"
)

var UID = os.Getenv("CENSYS_UID")
var SECRET = os.Getenv("CENSYS_SECRET")

var api = NewCensysAPI(UID, SECRET, nil)

func Test_Credentials(t *testing.T) {
	if UID == "" || SECRET == "" {
		t.Fatal("UID and/or SECRET must not empty.")
	}
}
