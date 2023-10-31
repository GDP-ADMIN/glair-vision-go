// Package ocr is collection of functions that interacts
// with GLAIR Vision OCR products
//
// This package is not meant to be used separately without the client
package ocr

import (
	"context"
	"os"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/internal"
)

// OCR provides functions to interact with GLAIR Vision
// OCR products
type OCR struct {
	config *glair.Config
}

// OCRResult is wrapper object for OCR API responses
type OCRResult[T any] struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Read   T      `json:"read"`
}

// OCRImage stores image data from OCR API response
type OCRImage struct {
	Photo string `json:"photo"`
	Sign  string `json:"sign"`
}

// OCRField stores field information from the provided image
type OCRField struct {
	Confidence int    `json:"confidence"`
	Value      string `json:"value"`
}

// OCRQualities stores image quality information from the provided image
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
func New(config *glair.Config) *OCR {
	return &OCR{
		config: config,
	}
}

func (ocr *OCR) readFile(file interface{}) (*os.File, error) {
	var input *os.File

	switch param := file.(type) {
	case string:
		file, err := os.Open(param)
		if err != nil {
			return nil, &glair.Error{
				Code:    glair.ErrorCodeInvalidFile,
				Message: "cannot read file from path",
				Err:     err,
			}
		}

		input = file
	case *os.File:
		input = file.(*os.File)
	default:
		return nil, &glair.Error{
			Code:    glair.ErrorCodeInvalidFile,
			Message: "invalid file type. Supported types are string or *os.File",
		}
	}

	return input, nil
}

// Ktp
func (ocr *OCR) Ktp(
	ctx context.Context,
	file interface{},
) (KTP, error) {
	ktp, err := ocr.readFile(file)
	if err != nil {
		return KTP{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "ktp")

	return internal.MakeRequest[KTP](ctx, url, ocr.config, ktp)
}

// KtpWithQuality
func (ocr *OCR) KTPWithQuality(
	ctx context.Context,
	file interface{},
) (KTPWithQuality, error) {
	ktp, err := ocr.readFile(file)
	if err != nil {
		return KTPWithQuality{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "ktp/quality")

	return internal.MakeRequest[KTPWithQuality](ctx, url, ocr.config, ktp)
}

func (ocr *OCR) NPWP(
	ctx context.Context,
	file interface{},
) (NPWP, error) {
	npwp, err := ocr.readFile(file)
	if err != nil {
		return NPWP{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "npwp")

	return internal.MakeRequest[NPWP](ctx, url, ocr.config, npwp)
}

func (ocr *OCR) KK(
	ctx context.Context,
	file interface{},
) (KK, error) {
	kk, err := ocr.readFile(file)
	if err != nil {
		return KK{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "kk")

	return internal.MakeRequest[KK](ctx, url, ocr.config, kk)
}

func (ocr *OCR) STNK(
	ctx context.Context,
	file interface{},
) (STNK, error) {
	stnk, err := ocr.readFile(file)
	if err != nil {
		return STNK{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "stnk")

	return internal.MakeRequest[STNK](ctx, url, ocr.config, stnk)
}

func (ocr *OCR) BPKB(
	ctx context.Context,
	file interface{},
) (BPKB, error) {
	bpkb, err := ocr.readFile(file)
	if err != nil {
		return BPKB{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "bpkb")

	return internal.MakeRequest[BPKB](ctx, url, ocr.config, bpkb)
}
