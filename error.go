package glair

// ErrorCode is unique code to quickly distinguish errors from
// GLAIR Vision SDK
type ErrorCode string

const (
	// ErrorCodeInvalidFile is returned when the SDK fails
	// to read the input
	ErrorCodeInvalidFile ErrorCode = "INVALID_FILE"
	// ErrorCodeFileCorrupted is returned when the provided file
	// is corrupted or too low on quality
	ErrorCodeFileCorrupted ErrorCode = "FILE_CORRUPTED"
	// ErrorCodeInvalidURL is returned when user provides
	// an invalid base URL in the configuration object
	ErrorCodeInvalidURL ErrorCode = "INVAlID_URL"
	// ErrorCodeBadClient is returned when the provided HTTP
	// client unable to send HTTP request to GLAIR Vision API
	ErrorCodeBadClient ErrorCode = "BAD_CLIENT"
	// ErrorCodeInvalidResponse is returned when GLAIR Vision API
	// returns an unexpected response
	//
	// Genrally, this error is impossible to be returned.
	// If you encounter this error, please contact us
	ErrorCodeInvalidResponse ErrorCode = "INVALID_RESPONSE"
)

// ResponseBody represents the response body returned
// by GLAIR Vision API if the response body exists
type ResponseBody struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
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

	// Body represents response returned by GLAIR Vision API
	// if available
	Body ResponseBody
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}
