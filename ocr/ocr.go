// Package ocr is collection of functions that interacts
// with GLAIR Vision OCR products
//
// This package is not meant to be used separately without the client
package ocr

import (
	"github.com/glair-ai/glair-vision-go/config"
)

// OCR provides functions to interact with GLAIR Vision
// OCR products
type OCR struct {
	config config.Config
}

// OCRResult is wrapper object for OCR API responses
type OCRResult[T any] struct {
	Status string
	Reason string
	Data   T
}

// OCRImage stores image data from OCR API response
type OCRImage struct {
	Photo string
	Sign  string
}

// OCRField stores field information from the provided image
type OCRField struct {
	Confidence int
	Value      string
}

type OCRQualities struct {
	IsBlurred bool `json:"is_blurred"`
	IsBright  bool `json:"is_bright"`
	IsCopy    bool `json:"is_copy"`
	IsDark    bool `json:"is_dark"`
	IsFlash   bool `json:"is_flash"`
	IsKtp     bool `json:"is_ktp"`
	IsRotated bool `json:"is_rotated"`
}

// New creates a GLAIR Vision OCR API Client from the provided config
func New(config config.Config) *OCR {
	return &OCR{
		config: config,
	}
}
