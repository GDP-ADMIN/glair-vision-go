package config

import (
	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func NewClient() *client.Client {
	config := glair.NewConfig(
		"",
		"",
		"",
	)
	config = config.WithBaseURL("https://api.vision.glair.ai")
	return client.New(config)
}
