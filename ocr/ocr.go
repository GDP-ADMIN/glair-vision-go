// Package ocr is collection of functions that interacts
// with GLAIR Vision OCR products
//
// This package is not meant to be used separately without the client
package ocr

import "github.com/glair-ai/glair-vision-go/config"

// OCR provides functions to interact with GLAIR Vision
// OCR products
type OCR struct {
	config config.Config
}

func New(config config.Config) *OCR {
	return &OCR{
		config: config,
	}
}
