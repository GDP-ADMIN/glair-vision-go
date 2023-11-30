// Package client provides API client that can be used
// to interact with GLAIR Vision products
//
// Generally, this is the only package that you need to
// interact with GLAIR Vision API outside the needed
// definitions
package client

import (
	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/face"
	"github.com/glair-ai/glair-vision-go/identity"
	"github.com/glair-ai/glair-vision-go/ocr"
)

// Client provides API to interact with GLAIR Vision products
type Client struct {
	// Config provides basic configurations that are used to interact when
	// calling the GLAIR Vision API
	Config *glair.Config
	// Ocr provides API interface to interact with GLAIR Vision OCR API
	Ocr *ocr.OCR
	// FaceBio provides API interface to interact with GLAIR Vision face
	// biometrics API
	FaceBio *face.FaceBio
	// FaceBio provides API interface to interact with GLAIR Vision
	// identity verification API
	Identity *identity.Identity
}

// New instatiates a client instance with the provided configuration
// and return a pointer to the new client
func New(config *glair.Config) *Client {
	return &Client{
		Config: config,

		Ocr:      ocr.New(config),
		FaceBio:  face.New(config),
		Identity: identity.New(config),
	}
}
