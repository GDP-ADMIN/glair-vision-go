package glair

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	e := &Error{
		Message: "something went wrong",
		Err:     errors.New("underlying cause"),
	}

	assert.Equal(t, "something went wrong", e.Error())
}

func TestError_Unwrap(t *testing.T) {
	underlying := errors.New("underlying cause")
	e := &Error{
		Message: "something went wrong",
		Err:     underlying,
	}

	assert.Equal(t, underlying, e.Unwrap())
}
