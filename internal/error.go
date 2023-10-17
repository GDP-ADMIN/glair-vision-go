package internal

type responseBody struct {
	Status string
	Reason string
}

type RequestError struct {
	StatusCode int
	Body       responseBody
	Err        string
}
