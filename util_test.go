package glair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	str := "Hello world!"
	expected := &str

	assert.Equal(t, expected, String(str))
}
