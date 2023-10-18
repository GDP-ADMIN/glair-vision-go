package client

import (
	"testing"

	"github.com/glair-ai/glair-vision-go/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotPanics(t, func() {
		config := config.New("username", "password", "")

		New(config)
	})
}
