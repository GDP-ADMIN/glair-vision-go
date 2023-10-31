package glair

type ErrorCode string

const (
	ErrorCodeInvalidFile     ErrorCode = "INVALID_FILE"
	ErrorCodeFileCorrupted   ErrorCode = "FILE_CORRUPTED"
	ErrorCodeInvalidURL      ErrorCode = "INVAlID_URL"
	ErrorCodeBadClient       ErrorCode = "BAD_CLIENT"
	ErrorCodeInvalidResponse ErrorCode = "INVALID_RESPONSE"
)

type ResponseBody struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type Error struct {
	Code    ErrorCode
	Message string
	Err     error

	Body ResponseBody
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}
