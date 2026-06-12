// Package ocr is a collection of functions and objects that interacts
// with GLAIR Vision OCR products and its results
package ocr

import (
	"context"

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

// OCRField is a unified generic type that stores field information
// of OCR result from the given file
type OCRField[T any] struct {
	Confidence     *float32 `json:"confidence,omitempty"`
	ConfidenceText *float32 `json:"confidence_text,omitempty"`
	Polygon        [][]int  `json:"polygon,omitempty"`
	PageIndex      *int     `json:"page_index,omitempty"`
	ValueOriginal  *string  `json:"value_original,omitempty"`
	Value          *T       `json:"value,omitempty"`
}

// OCRStringField is a type alias for OCRField with string value
type OCRStringField = OCRField[string]

// OCRIntField is a type alias for OCRField with int64 value
type OCRIntField = OCRField[int64]

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

// New creates a GLAIR Vision OCR API Client with
// the given config
func New(config *glair.Config) *OCR {
	return &OCR{
		config: config,
	}
}

// recognize is a helper for standard multipart OCR endpoints.
func (ocr *OCR) recognize(
	ctx context.Context,
	endpoint string,
	input glair.OCRInput,
	target any,
) error {
	file, err := internal.ReadFile(input.Image)
	if err != nil {
		return err
	}

	params := internal.RequestParameters{
		Url:       ocr.config.GetEndpointURL("ocr", endpoint),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": file,
		},
	}

	return internal.MakeMultipartRequest(
		ctx,
		params,
		ocr.config,
		target,
	)
}

// KTP performs OCR on the given file using KTP model.
func (ocr *OCR) KTP(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "ktp", input, target)
}

// KTPWithQualities performs OCR on the given file using KTP model
// and supplements it with file quality data.
func (ocr *OCR) KTPWithQualities(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "ktp-with-qualities", input, target)
}

// NPWP performs OCR on the given file using NPWP model.
func (ocr *OCR) NPWP(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "npwp", input, target)
}

// KK performs OCR on the given file using Kartu Keluarga model.
func (ocr *OCR) KK(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "kk", input, target)
}

// STNK performs OCR on the given file using STNK model.
func (ocr *OCR) STNK(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "stnk", input, target)
}

// SIM performs OCR on the given file using SIM model.
func (ocr *OCR) SIM(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "sim", input, target)
}

// Passport performs OCR on the given file using Passport model.
func (ocr *OCR) Passport(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "passport", input, target)
}

// Plate performs OCR on the given file using License Plate model.
func (ocr *OCR) Plate(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "plate", input, target)
}

// GeneralDocument performs OCR on the given file using all-purpose Document model.
func (ocr *OCR) GeneralDocument(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "general-document", input, target)
}

// Invoice performs OCR on the given file using Invoice model.
func (ocr *OCR) Invoice(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "invoice", input, target)
}

// Receipt performs OCR on the given file using Receipt model.
func (ocr *OCR) Receipt(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "receipt", input, target)
}

// BankStatement performs OCR on the given file using Bank Statement model.
func (ocr *OCR) BankStatement(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "bank-statement", input, target)
}

// SKPR performs OCR on the given file using SKPR model.
func (ocr *OCR) SKPR(
	ctx context.Context,
	input glair.OCRInput,
	target any,
) error {
	return ocr.recognize(ctx, "skpr", input, target)
}

// BPKB performs OCR on the given file using BPKB model.
func (ocr *OCR) BPKB(
	ctx context.Context,
	input glair.BPKBInput,
	target any,
) error {
	file, err := internal.ReadFile(input.Image)
	if err != nil {
		return err
	}

	params := internal.RequestParameters{
		Url:       ocr.config.GetEndpointURL("ocr", "bpkb"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": file,
			"page":  input.Page,
		},
	}

	return internal.MakeMultipartRequest(
		ctx,
		params,
		ocr.config,
		target,
	)
}

// KTPSessions sends session request for passive liveness
// using the prebuilt web page.
func (ocr *OCR) KTPSessions(
	ctx context.Context,
	input glair.SessionsInput,
	target any,
) error {
	payload := map[string]any{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}

	params := internal.RequestParameters{
		Url:  ocr.config.GetEndpointURL("ocr", "ktp-sessions"),
		Body: payload,
	}

	return internal.MakeJSONRequest(
		ctx,
		params,
		ocr.config,
		target,
	)
}

// NPWPSessions sends session request for passive liveness
// using the prebuilt web page.
func (ocr *OCR) NPWPSessions(
	ctx context.Context,
	input glair.SessionsInput,
	target any,
) error {
	payload := map[string]any{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}

	params := internal.RequestParameters{
		Url:  ocr.config.GetEndpointURL("ocr", "npwp-sessions"),
		Body: payload,
	}

	return internal.MakeJSONRequest(
		ctx,
		params,
		ocr.config,
		target,
	)
}
