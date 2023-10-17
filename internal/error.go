package internal

import "fmt"

type responseBody struct {
	Status string
	Reason string
}

type RequestError struct {
	StatusCode int
	Body       responseBody
}

func (e RequestError) Error() string {
	return fmt.Sprintf("Failed to process request: %s", e.Body.Reason)
}
