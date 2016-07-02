package gocensys

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// BaseURL represents the Censys version 1 API URL.
	BaseURL = "https://www.censys.io/api/v1"
)

// CensysAPI represents a censys API client used for accessing available endpoints.
type CensysAPI struct {
	// HTTPClient is the HTTP client that will be used in the API requests.
	HTTPClient *http.Client

	// UID represents your API key.
	UID string

	// Secret is required for accessing the API.
	Secret string
}

// NewCensysAPI takes user specific API ID (uid), secret and returns a CensysAPI struct for that user.
func NewCensysAPI(uid, secret string, httpClient *http.Client) *CensysAPI {
	return &CensysAPI{
		HTTPClient: http.DefaultClient,
		UID:        uid,
		Secret:     secret,
	}
}

// apiGet performs "GET" requests on the specified endpoint.
func (c CensysAPI) apiGet(endpoint string, form url.Values, data interface{}) error {
	req, err := http.NewRequest("GET", BaseURL+endpoint, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.UID, c.Secret)
	req.Form = form

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return decodeResponse(resp, data)
}

// apiPost performs "POST" requests on the specified endpoint.
func (c CensysAPI) apiPost(endpoint string, query map[string]interface{}, form url.Values, data interface{}) error {
	jsonStr, err := json.Marshal(query)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", BaseURL+endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(c.UID, c.Secret)
	req.Form = form

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return decodeResponse(resp, data)
}

// decodeResponse unmarshals the a json file into a Go struct
func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != http.StatusOK {
		return newAPIError(resp)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
