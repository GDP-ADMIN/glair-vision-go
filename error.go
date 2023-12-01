package glair

// ErrorCode is unique code to quickly distinguish errors from
// GLAIR Vision SDK
type ErrorCode string

const (
	// ErrorCodeFileError is returned when the SDK fails
	// to read or parse the given input file
	ErrorCodeFileError ErrorCode = "FILE_ERROR"
	// ErrorCodeNetworkError is returned when the SDK is unable to
	// complete the HTTP request to GLAIR Vision API with the given
	// HTTP client
	ErrorCodeNetworkError ErrorCode = "NETWORK_ERROR"
	// ErrorCodeTimeout is returned when the HTTP request sent
	// by the SDK has timed out
	//
	// To solve this problem, you can increase the timeout
	// value from the context
	ErrorCodeTimeout ErrorCode = "TIMEOUT"
	// ErrorCodeForbidden is returned when the provided credentials
	// have insufficient access rights to the requested endpoint
	//
	// If you think this is a mistake, please contact us
	ErrorCodeForbidden ErrorCode = "FORBIDDEN"
	// ErrorCodeAPIError is returned when GLAIR Vision API
	// returns a non-OK response. In this case, please
	// check the Body property for more details on the error
	ErrorCodeAPIError ErrorCode = "API_ERROR"
	// ErrorCodeInvalidResponse is returned when GLAIR Vision API
	// returns an unexpected response
	//
	// Generally, this error is impossible to be returned.
	// If you encounter this error, please contact us
	ErrorCodeInvalidResponse ErrorCode = "INVALID_RESPONSE"
)

// Response represents the response returned
// by GLAIR Vision API if request returned an error
type Response struct {
	Status int                    `json:"status,omitempty"`
	Body   map[string]interface{} `json:"body"`
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
