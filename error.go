package glair

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidBaseUrl = errors.New("invalid base URL configuration")
	ErrFileRequired   = errors.New("input file is required")
	ErrInvalidFile    = errors.New("input file is corrupted")
	ErrBadClient      = errors.New("invalid http client")

	// this should not happen

	ErrInvalidResponseBody = errors.New("invalid response body format")
	ErrInvalidFormData     = errors.New("failed to create form data")
)

type responseBody struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

// RequestError is an error that occurs when GLAIR Vision
// API does not return HTTP Status OK
type RequestError struct {
	StatusCode   int    `json:"status_code"`
	ResponseBody []byte `json:"body"`
}

func (e RequestError) Error() string {
	var body responseBody
	json.Unmarshal(e.ResponseBody, &body)

	return body.Reason
}

// Body returns raw response body
func (e RequestError) Body() responseBody {
	var body responseBody
	json.Unmarshal(e.ResponseBody, &body)

	return body
}
