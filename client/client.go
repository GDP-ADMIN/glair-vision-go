// Package client is a collection of objects that provides
// API to interact with GLAIR Vision products
package client

import (
	"github.com/glair-ai/glair-vision-go/config"
	"github.com/glair-ai/glair-vision-go/ocr"
)

// Client provides API to interact with GLAIR Vision products
type Client struct {
	Ocr *ocr.OCR
}

// New instatiates a client instance with the provided configuration
// and return a pointer to the new client
func New(config *config.Config) *Client {
	return &Client{
		Ocr: ocr.New(config),
	}
}
