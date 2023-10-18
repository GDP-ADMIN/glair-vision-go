package internal

import "fmt"

type responseBody struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type RequestError struct {
	StatusCode int
	Body       responseBody
}

func (e RequestError) Error() string {
	return fmt.Sprintf("Failed to process request. status: %d. reason: %s", e.StatusCode, e.Body.Reason)
}
