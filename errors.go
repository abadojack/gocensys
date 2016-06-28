package gocensys

import (
	"fmt"
	"net/http"
)

// APIError represents an error reponse
type APIError struct {
	ErrorCode int
	ErrorStr  string
	Request   *http.Request
}

// ApiError supports the error interface
func (apiErr APIError) Error() string {
	return fmt.Sprintf("%s %s returned status %d, %s", apiErr.Request.Method, apiErr.Request.URL, apiErr.ErrorCode, apiErr.ErrorStr)
}

func newAPIError(resp *http.Response) *APIError {
	return &APIError{
		ErrorCode: resp.StatusCode,
		ErrorStr:  resp.Status,
		Request:   resp.Request,
	}
}
