package client

import (
	"testing"

	"github.com/glair-ai/glair-vision-go/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	config := config.New("username", "password", "")
	client := New(config)

	assert.Equal(t, config, client.config)
}
