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

// OCRField is stores field information of OCR result from the given file
type OCRField struct {
	// Confidence represents OCR accuracy of the value
	Confidence float32 `json:"confidence"`

	// Polygon represents value position from the given image
	//
	// Do note that not all OCR API provides this value
	Polygon [][]int `json:"polygon"`
}

// OCRStringField stores field information that can be represented
// as a string of OCR result from the given file
type OCRStringField struct {
	OCRField
	Value string `json:"value"`
}

// OCRIntField stores field information that can be represented
// as an integer of OCR result from the given file
type OCRIntField struct {
	OCRField
	Value int64 `json:"value"`
}

// PaperlessOCRField stores field information of OCR result from the
// given file in addition of page index and original text data
type PaperlessOCRField struct {
	OCRField

	ConfidenceText float32 `json:"confidence_text,omitempty"`
	PageIndex      int     `json:"page_index"`
	ValueOriginal  string  `json:"value_original,omitempty"`
}

// PaperlessOCRStringField stores field information of OCR result from the
// given file that can be represented as a string in addition of page
// index and original text data
type PaperlessOCRStringField struct {
	PaperlessOCRField

	Value string `json:"value,omitempty"`
}

// PaperlessOCRIntField stores field information of OCR result from the
// given file that can be represented as a string in addition of page
// index and original text data
type PaperlessOCRIntField struct {
	PaperlessOCRField

	Value int64 `json:"value,omitempty"`
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

// New creates a GLAIR Vision OCR API Client with
// the given config
func New(config *glair.Config) *OCR {
	return &OCR{
		config: config,
	}
}

// Ktp performs OCR on the given file using KTP model
//
// API Docs: https://docs.glair.ai/vision/ktp
func (ocr *OCR) Ktp(
	ctx context.Context,
	input glair.OCRInput,
) (KTP, error) {
	ktp, err := internal.ReadFile(input.Image)
	if err != nil {
		return KTP{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "ktp")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": ktp,
		},
	}

	return internal.MakeMultipartRequest[KTP](ctx, params, ocr.config)
}

// KtpWithQuality performs OCR on the given file using KTP model
// and supplements it with file quality data
//
// API Docs: https://docs.glair.ai/vision/ktp
func (ocr *OCR) KtpWithQuality(
	ctx context.Context,
	input glair.OCRInput,
) (KTPWithQuality, error) {
	ktp, err := internal.ReadFile(input.Image)
	if err != nil {
		return KTPWithQuality{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "ktp/qualities")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": ktp,
		},
	}

	return internal.MakeMultipartRequest[KTPWithQuality](ctx, params, ocr.config)
}

// NPWP performs OCR on the given file using NPWP model
//
// API Docs: https://docs.glair.ai/vision/npwp
func (ocr *OCR) NPWP(
	ctx context.Context,
	input glair.OCRInput,
) (NPWP, error) {
	npwp, err := internal.ReadFile(input.Image)
	if err != nil {
		return NPWP{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "npwp")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": npwp,
		},
	}

	return internal.MakeMultipartRequest[NPWP](ctx, params, ocr.config)
}

// KK performs OCR on the given file using Kartu Keluarga model
//
// API Docs: https://docs.glair.ai/vision/kk
func (ocr *OCR) KK(
	ctx context.Context,
	input glair.OCRInput,
) (KK, error) {
	kk, err := internal.ReadFile(input.Image)
	if err != nil {
		return KK{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "kk")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": kk,
		},
	}

	return internal.MakeMultipartRequest[KK](ctx, params, ocr.config)
}

// STNK performs OCR on the given file using STNK model
//
// API Docs: https://docs.glair.ai/vision/stnk
func (ocr *OCR) STNK(
	ctx context.Context,
	input glair.OCRInput,
) (STNK, error) {
	stnk, err := internal.ReadFile(input.Image)
	if err != nil {
		return STNK{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "stnk")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": stnk,
		},
	}

	return internal.MakeMultipartRequest[STNK](ctx, params, ocr.config)
}

// BPKB performs OCR on the given file using BPKB model
//
// API Docs: https://docs.glair.ai/vision/bpkb
func (ocr *OCR) BPKB(
	ctx context.Context,
	input glair.BPKBInput,
) (BPKB, error) {
	bpkb, err := internal.ReadFile(input.Image)
	if err != nil {
		return BPKB{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "bpkb")

	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": bpkb,
			"page":  input.Page,
		},
	}

	return internal.MakeMultipartRequest[BPKB](ctx, params, ocr.config)
}

// Passport performs OCR on the given file using Passport model
//
// API Docs: https://docs.glair.ai/vision/passport
func (ocr *OCR) Passport(
	ctx context.Context,
	input glair.OCRInput,
) (Passport, error) {
	passport, err := internal.ReadFile(input.Image)
	if err != nil {
		return Passport{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "passport")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": passport,
		},
	}

	return internal.MakeMultipartRequest[Passport](ctx, params, ocr.config)
}

// Plate performs OCR on the given file using License Plate model
//
// API Docs: https://docs.glair.ai/vision/plate
func (ocr *OCR) Plate(
	ctx context.Context,
	input glair.OCRInput,
) (Plate, error) {
	plate, err := internal.ReadFile(input.Image)
	if err != nil {
		return Plate{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "plate")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": plate,
		},
	}

	return internal.MakeMultipartRequest[Plate](ctx, params, ocr.config)
}

// GeneralDocument performs OCR on the given file using all-purpose Document model
//
// API Docs: https://docs.glair.ai/vision/general-document
func (ocr *OCR) GeneralDocument(
	ctx context.Context,
	input glair.OCRInput,
) (GeneralDocument, error) {
	generalDocument, err := internal.ReadFile(input.Image)
	if err != nil {
		return GeneralDocument{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "general-document")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": generalDocument,
		},
	}

	return internal.MakeMultipartRequest[GeneralDocument](ctx, params, ocr.config)
}

// Invoice performs OCR on the given file using Invoice model
//
// API Docs: https://docs.glair.ai/vision/invoice
func (ocr *OCR) Invoice(
	ctx context.Context,
	input glair.OCRInput,
) (Invoice, error) {
	invoice, err := internal.ReadFile(input.Image)
	if err != nil {
		return Invoice{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "invoice")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": invoice,
		},
	}

	return internal.MakeMultipartRequest[Invoice](ctx, params, ocr.config)
}

// Receipt performs OCR on the given file using Receipt model
//
// API Docs: https://docs.glair.ai/vision/receipt
func (ocr *OCR) Receipt(
	ctx context.Context,
	input glair.OCRInput,
) (Receipt, error) {
	receipt, err := internal.ReadFile(input.Image)
	if err != nil {
		return Receipt{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "receipt")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": receipt,
		},
	}

	return internal.MakeMultipartRequest[Receipt](ctx, params, ocr.config)
}

// BankStatement performs OCR on the given file using Bank Statement model
//
// API Docs: https://docs.glair.ai/vision/bank-statement
func (ocr *OCR) BankStatement(
	ctx context.Context,
	input glair.OCRInput,
) (BankStatement, error) {
	bankStatement, err := internal.ReadFile(input.Image)
	if err != nil {
		return BankStatement{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "bank-statement")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": bankStatement,
		},
	}

	return internal.MakeMultipartRequest[BankStatement](ctx, params, ocr.config)
}

// SKPR performs OCR on the given file using SKPR model
//
// API Docs: https://docs.glair.ai/vision/skpr
func (ocr *OCR) SKPR(
	ctx context.Context,
	input glair.OCRInput,
) (SKPR, error) {
	skpr, err := internal.ReadFile(input.Image)
	if err != nil {
		return SKPR{}, err
	}

	url := ocr.config.GetEndpointURL("ocr", "skpr")
	params := internal.RequestParameters{
		Url:       url,
		RequestID: input.RequestID,
		Body: map[string]interface{}{
			"image": skpr,
		},
	}

	return internal.MakeMultipartRequest[SKPR](ctx, params, ocr.config)
}

// KtpSessions sends session request for passive liveness
// using the prebuilt web page
//
// API Docs: https://docs.glair.ai/vision/ktp-sessions
func (ocr *OCR) KtpSessions(
	ctx context.Context,
	input glair.SessionsInput,
) (glair.Session, error) {
	payload := map[string]interface{}{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}
	url := ocr.config.GetEndpointURL("ocr", "ktp-sessions")
	params := internal.RequestParameters{
		Url:  url,
		Body: payload,
	}

	return internal.MakeJSONRequest[glair.Session](ctx, params, ocr.config)
}

// NPWPSessions sends session request for passive liveness
// using the prebuilt web page
//
// API Docs: https://docs.glair.ai/vision/npwp-sessions
func (ocr *OCR) NPWPSessions(
	ctx context.Context,
	input glair.SessionsInput,
) (glair.Session, error) {
	payload := map[string]interface{}{
		"success_url": input.SuccessURL,
	}

	if input.CancelURL != nil {
		payload["cancel_url"] = input.CancelURL
	}
	url := ocr.config.GetEndpointURL("ocr", "npwp-sessions")
	params := internal.RequestParameters{
		Url:  url,
		Body: payload,
	}

	return internal.MakeJSONRequest[glair.Session](ctx, params, ocr.config)
}
