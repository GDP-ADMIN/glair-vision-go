// Package ocr is a collection of functions and objects that interacts
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

// OCRField stores field information from OCR result of the given file
type OCRField struct {
	// Confidence represents OCR accuracy of the value
	Confidence float32 `json:"confidence,omitempty"`
	Value      string  `json:"value,omitempty"`
	// Polygon represents value position from the given image
	//
	// Do note that not all OCR API provides this value
	Polygon [][]int `json:"polygon,omitempty"`
}

// OCRQualities stores image quality information from OCR result of
// the given file
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
				Message: "Cannot read file from the given path.",
				Err:     err,
			}
		}

		input = file
	case *os.File:
		input = file.(*os.File)
	default:
		return nil, &glair.Error{
			Code:    glair.ErrorCodeInvalidFile,
			Message: "Invalid file type is provided. Valid file types are string or *os.File",
		}
	}

	return input, nil
}

// Ktp performs OCR on the given file using KTP model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/ktp
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

// KtpWithQuality performs OCR on the given file using KTP model
// and supplements it with file quality data
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/ktp
func (ocr *OCR) KtpWithQuality(
	ctx context.Context,
	file interface{},
) (KTPWithQuality, error) {
	ktp, err := ocr.readFile(file)
	if err != nil {
		return KTPWithQuality{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "ktp/qualities")

	return internal.MakeRequest[KTPWithQuality](ctx, url, ocr.config, ktp)
}

// NPWP performs OCR on the given file using NPWP model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/npwp
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

// KK performs OCR on the given file using Kartu Keluarga model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/kk
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

// STNK performs OCR on the given file using STNK model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/stnk
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

// BPKB performs OCR on the given file using BPKB model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/bpkb
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

// Passport performs OCR on the given file using Passport model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/passport
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

// Plate performs OCR on the given file using License Plate model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/plate
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

// GeneralDocument performs OCR on the given file using all-purpose Document model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/general-document
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

// Invoice performs OCR on the given file using Invoice model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/invoice
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

// Receipt performs OCR on the given file using Receipt model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/receipt
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

// BankStatement performs OCR on the given file using Bank Statement model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/bank-statement
func (ocr *OCR) BankStatement(
	ctx context.Context,
	file interface{},
) (BankStatement, error) {
	bankStatement, err := ocr.readFile(file)
	if err != nil {
		return BankStatement{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "bank-statement")

	return internal.MakeRequest[BankStatement](ctx, url, ocr.config, bankStatement)
}

// SKPR performs OCR on the given file using SKPR model
//
// File must be provided as a string that represents a path or
// an interface of *os.File
//
// API Docs: https://docs.glair.ai/vision/skpr
func (ocr *OCR) SKPR(
	ctx context.Context,
	file interface{},
) (SKPR, error) {
	skpr, err := ocr.readFile(file)
	if err != nil {
		return SKPR{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "skpr")

	return internal.MakeRequest[SKPR](ctx, url, ocr.config, skpr)
}
