package internal

import (
	"fmt"
	"os"

	"github.com/glair-ai/glair-vision-go"
)

// ReadFile reads file from a string representation or
// an actual file
func ReadFile(file interface{}) (*os.File, error) {
	var input *os.File

	switch param := file.(type) {
	case string:
		file, err := os.Open(param)
		if err != nil {
			return nil, &glair.Error{
				Code:    glair.ErrorCodeInvalidFile,
				Message: fmt.Sprintf("Cannot read file from path %s", param),
				Err:     err,
			}
		}

		input = file
	case *os.File:
		input = file.(*os.File)
	default:
		return nil, &glair.Error{
			Code:    glair.ErrorCodeInvalidFile,
			Message: "Invalid file type is provided. Valid file types are string or *os.File",
		}
	}

	return input, nil
}
