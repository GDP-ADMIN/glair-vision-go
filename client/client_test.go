package client

import (
	"testing"

	"github.com/glair-ai/glair-vision-go"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	config := glair.NewConfig("username", "password", "api_key")

	client := New(config)

	assert.Equal(t, config, client.Config)
}
