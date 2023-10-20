// Package client provides API client that can be used
// to interact with GLAIR Vision products
package client

import (
	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/ocr"
)

// Client provides API to interact with GLAIR Vision products
type Client struct {
	Config *glair.Config

	Ocr *ocr.OCR
}

// New instatiates a client instance with the provided configuration
// and return a pointer to the new client
func New(config *glair.Config) *Client {
	return &Client{
		Config: config,
		Ocr:    ocr.New(config),
	}
}
