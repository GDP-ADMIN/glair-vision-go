// Package ocr is collection of functions and objects that interacts
// with GLAIR Vision OCR products and its results
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
	Photo string `json:"photo,omitempty"`
	Sign  string `json:"sign,omitempty"`
}

// OCRField stores field information from the provided image
type OCRField struct {
	Confidence int      `json:"confidence,omitempty"`
	Value      string   `json:"value,omitempty"`
	Polygon    [][2]int `json:"polygon,omitempty"`
}

// OCRQualities stores image quality information from the provided image
type OCRQualities struct {
	IsBlurred bool `json:"is_blurred,omitempty"`
	IsBright  bool `json:"is_bright,omitempty"`
	IsCopy    bool `json:"is_copy,omitempty"`
	IsDark    bool `json:"is_dark,omitempty"`
	IsFlash   bool `json:"is_flash,omitempty"`
	IsKtp     bool `json:"is_ktp,omitempty"`
	IsRotated bool `json:"is_rotated,omitempty"`
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
func (ocr *OCR) KtpWithQuality(
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

func (ocr *OCR) Passport(
	ctx context.Context,
	file interface{},
) (Passport, error) {
	passport, err := ocr.readFile(file)
	if err != nil {
		return Passport{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "passport")

	return internal.MakeRequest[Passport](ctx, url, ocr.config, passport)
}

func (ocr *OCR) Plate(
	ctx context.Context,
	file interface{},
) (Plate, error) {
	passport, err := ocr.readFile(file)
	if err != nil {
		return Plate{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "plate")

	return internal.MakeRequest[Plate](ctx, url, ocr.config, passport)
}

func (ocr *OCR) GeneralDocument(
	ctx context.Context,
	file interface{},
) (GeneralDocument, error) {
	generalDocument, err := ocr.readFile(file)
	if err != nil {
		return GeneralDocument{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "general-document")

	return internal.MakeRequest[GeneralDocument](ctx, url, ocr.config, generalDocument)
}

func (ocr *OCR) Invoice(
	ctx context.Context,
	file interface{},
) (Invoice, error) {
	invoice, err := ocr.readFile(file)
	if err != nil {
		return Invoice{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "invoice")

	return internal.MakeRequest[Invoice](ctx, url, ocr.config, invoice)
}

func (ocr *OCR) Receipt(
	ctx context.Context,
	file interface{},
) (Receipt, error) {
	receipt, err := ocr.readFile(file)
	if err != nil {
		return Receipt{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "receipt")

	return internal.MakeRequest[Receipt](ctx, url, ocr.config, receipt)
}
