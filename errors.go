package gocensys

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIError represents an error response
type APIError struct {
	ErrorCode int
	ErrorStr  string
	Request   *http.Request
}

// ApiError supports the error interface
func (apiErr APIError) Error() string {
	return fmt.Sprintf("%s %s returned %s", apiErr.Request.Method, apiErr.Request.URL, apiErr.ErrorStr)
}

func newAPIError(resp *http.Response) *APIError {
	respBody, _ := ioutil.ReadAll(resp.Body) // TODO : Please don't ignore this error.
	return &APIError{
		ErrorCode: resp.StatusCode,
		ErrorStr:  string(respBody),
		Request:   resp.Request,
	}
}
