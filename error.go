package glair

// ErrorCode is unique code to quickly distinguish errors from
// GLAIR Vision SDK
type ErrorCode string

const (
	// ErrorCodeInvalidFile is returned when the SDK fails
	// to read the input file
	ErrorCodeInvalidFile ErrorCode = "INVALID_FILE"
	// ErrorCodeFileCorrupted is returned when the provided file
	// is corrupted or too low on quality
	ErrorCodeFileCorrupted ErrorCode = "FILE_CORRUPTED"
	// ErrorCodeInvalidURL is returned when user provides
	// an invalid base URL in the configuration object
	ErrorCodeInvalidURL ErrorCode = "INVALID_URL"
	// ErrorCodeBadClient is returned when the provided HTTP
	// client unable to send HTTP request to GLAIR Vision API
	ErrorCodeBadClient ErrorCode = "BAD_CLIENT"
	// ErrorCodeAPIError is returned when GLAIR Vision API
	// returns a non-OK response. In this case, please
	// check the Body property for more details on the error
	ErrorCodeAPIError ErrorCode = "API_ERROR"
	// ErrorCodeInvalidResponse is returned when GLAIR Vision API
	// returns an unexpected response
	//
	// Genrally, this error is impossible to be returned.
	// If you encounter this error, please contact us
	ErrorCodeInvalidResponse ErrorCode = "INVALID_RESPONSE"
)

// Response represents the response returned
// by GLAIR Vision API if request returned an error
type Response struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// Error is an extended error object used by
// GLAIR Vision SDK
//
// Whenever error is returned, it is recommended
// to assert the error to this type
type Error struct {
	// Code represents unique code to quickly distinguish errors from
	// GLAIR Vision SDK
	Code ErrorCode
	// Message represents human-readable error object related
	// to the error
	Message string
	// Err is the original error object that is returned by the SDK
	Err error

	// Response represents response returned by GLAIR Vision API
	// if response is available
	Response Response
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}
