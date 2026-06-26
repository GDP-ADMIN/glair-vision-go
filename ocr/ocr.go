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

// PhotoField stores image data from OCR API response
type PhotoField struct {
	Confidence *float32 `json:"confidence,omitempty"`
	Polygon    [][]int  `json:"polygon,omitempty"`
	PageIndex  *int     `json:"page_index,omitempty"`
	Value      *string  `json:"value,omitempty"`
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

// QualityBooleanField stores boolean quality field information
type QualityBooleanField struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Value      *bool    `json:"value,omitempty"`
}

// BlurField stores blur quality field information
type BlurField struct {
	QualityBooleanField
	ConfidenceBlur *float64 `json:"confidence_blur,omitempty"`
}

// BrightField stores bright quality field information
type BrightField struct {
	QualityBooleanField
	ConfidenceBright *float64 `json:"confidence_bright,omitempty"`
}

// CropField stores crop quality field information
type CropField struct {
	QualityBooleanField
	ConfidenceCrop *float64 `json:"confidence_crop,omitempty"`
}

// DarkField stores dark quality field information
type DarkField struct {
	QualityBooleanField
	ConfidenceDark *float64 `json:"confidence_dark,omitempty"`
}

// FlashField stores flash quality field information
type FlashField struct {
	QualityBooleanField
	ConfidenceFlash *float64 `json:"confidence_flash,omitempty"`
}

// PhotocopyField stores photocopy quality field information
type PhotocopyField struct {
	QualityBooleanField
	ConfidencePhotocopy *float64 `json:"confidence_photocopy,omitempty"`
}

// ScreenField stores screen quality field information
type ScreenField struct {
	QualityBooleanField
	ConfidenceScreen *float64 `json:"confidence_screen,omitempty"`
}

// DocumentField stores document quality field information
type DocumentField struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Value      *string  `json:"value,omitempty"`
}

// RotateField stores rotation quality field information
type RotateField struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Value      *string  `json:"value,omitempty"`
	FixedImage *string  `json:"fixed_image,omitempty"`
}

// OCRQualitiesData stores image quality information from OCR result
type OCRQualitiesData struct {
	Blur      *BlurField      `json:"blur,omitempty"`
	Bright    *BrightField    `json:"bright,omitempty"`
	Crop      *CropField      `json:"crop,omitempty"`
	Dark      *DarkField      `json:"dark,omitempty"`
	Document  *DocumentField  `json:"document,omitempty"`
	Flash     *FlashField     `json:"flash,omitempty"`
	Photocopy *PhotocopyField `json:"photocopy,omitempty"`
	Rotate    *RotateField    `json:"rotate,omitempty"`
	Screen    *ScreenField    `json:"screen,omitempty"`
}

// New creates a GLAIR Vision OCR API Client with
// the given config
func New(config *glair.Config) *OCR {
	return &OCR{
		config: config,
	}
}

func recognize[T any](ctx context.Context, ocr *OCR, endpoint string, input glair.OCRInput) (T, error) {
	file, err := internal.ReadFile(input.Image)
	if err != nil {
		var zero T
		return zero, err
	}

	params := internal.RequestParameters{
		Url:       ocr.config.GetEndpointURL("ocr", endpoint),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": file,
		},
	}

	return internal.MakeMultipartRequest[T](ctx, params, ocr.config)
}

func recognizeRaw(ctx context.Context, ocr *OCR, endpoint string, input glair.OCRInput) ([]byte, error) {
	file, err := internal.ReadFile(input.Image)
	if err != nil {
		return nil, err
	}

	params := internal.RequestParameters{
		Url:       ocr.config.GetEndpointURL("ocr", endpoint),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": file,
		},
	}

	return internal.MakeMultipartRequestRaw(ctx, params, ocr.config)
}

// KTP performs OCR on the given file using KTP model
func (ocr *OCR) KTP(ctx context.Context, input glair.OCRInput) (KTP, error) {
	return recognize[KTP](ctx, ocr, "ktp", input)
}

// KTPRaw performs OCR on the given file using KTP model
// and returns the raw API response as bytes
func (ocr *OCR) KTPRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "ktp", input)
}

// KTPWithQuality performs OCR on the given file using KTP model
// and supplements it with file quality data
func (ocr *OCR) KTPWithQuality(ctx context.Context, input glair.OCRInput) (KTPWithQuality, error) {
	return recognize[KTPWithQuality](ctx, ocr, "ktp/qualities", input)
}

// KTPWithQualityRaw performs OCR on the given file using KTP model
// with quality supplements and returns the raw API response as bytes
func (ocr *OCR) KTPWithQualityRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "ktp/qualities", input)
}

// NPWP performs OCR on the given file using NPWP model
func (ocr *OCR) NPWP(ctx context.Context, input glair.OCRInput) (NPWP, error) {
	return recognize[NPWP](ctx, ocr, "npwp", input)
}

// NPWPRaw performs OCR on the given file using NPWP model
// and returns the raw API response as bytes
func (ocr *OCR) NPWPRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "npwp", input)
}

// KK performs OCR on the given file using KK model
func (ocr *OCR) KK(ctx context.Context, input glair.OCRInput) (KK, error) {
	return recognize[KK](ctx, ocr, "kk", input)
}

// KKRaw performs OCR on the given file using KK model
// and returns the raw API response as bytes
func (ocr *OCR) KKRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "kk", input)
}

// STNK performs OCR on the given file using STNK model
func (ocr *OCR) STNK(ctx context.Context, input glair.OCRInput) (STNK, error) {
	return recognize[STNK](ctx, ocr, "stnk", input)
}

// STNKRaw performs OCR on the given file using STNK model
// and returns the raw API response as bytes
func (ocr *OCR) STNKRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "stnk", input)
}

// SIM performs OCR on the given file using SIM model
func (ocr *OCR) SIM(ctx context.Context, input glair.OCRInput) (SIM, error) {
	return recognize[SIM](ctx, ocr, "sim", input)
}

// SIMRaw performs OCR on the given file using SIM model
// and returns the raw API response as bytes
func (ocr *OCR) SIMRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "sim", input)
}

// Passport performs OCR on the given file using Passport model
func (ocr *OCR) Passport(ctx context.Context, input glair.OCRInput) (Passport, error) {
	return recognize[Passport](ctx, ocr, "passport", input)
}

// PassportRaw performs OCR on the given file using Passport model
// and returns the raw API response as bytes
func (ocr *OCR) PassportRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "passport", input)
}

// Plate performs OCR on the given file using Plate model
func (ocr *OCR) Plate(ctx context.Context, input glair.OCRInput) (Plate, error) {
	return recognize[Plate](ctx, ocr, "plate", input)
}

// PlateRaw performs OCR on the given file using Plate model
// and returns the raw API response as bytes
func (ocr *OCR) PlateRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "plate", input)
}

// GeneralDocument performs OCR on the given file using General Document model
func (ocr *OCR) GeneralDocument(ctx context.Context, input glair.OCRInput) (GeneralDocument, error) {
	return recognize[GeneralDocument](ctx, ocr, "general-document", input)
}

// GeneralDocumentRaw performs OCR on the given file using General Document model
// and returns the raw API response as bytes
func (ocr *OCR) GeneralDocumentRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "general-document", input)
}

// Invoice performs OCR on the given file using Invoice model
func (ocr *OCR) Invoice(ctx context.Context, input glair.OCRInput) (Invoice, error) {
	return recognize[Invoice](ctx, ocr, "invoice", input)
}

// InvoiceRaw performs OCR on the given file using Invoice model
// and returns the raw API response as bytes
func (ocr *OCR) InvoiceRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "invoice", input)
}

// Receipt performs OCR on the given file using Receipt model
func (ocr *OCR) Receipt(ctx context.Context, input glair.OCRInput) (Receipt, error) {
	return recognize[Receipt](ctx, ocr, "receipt", input)
}

// ReceiptRaw performs OCR on the given file using Receipt model
// and returns the raw API response as bytes
func (ocr *OCR) ReceiptRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "receipt", input)
}

// BankStatement performs OCR on the given file using Bank Statement model
func (ocr *OCR) BankStatement(ctx context.Context, input glair.OCRInput) (BankStatement, error) {
	return recognize[BankStatement](ctx, ocr, "bank-statement", input)
}

// BankStatementRaw performs OCR on the given file using Bank Statement model
// and returns the raw API response as bytes
func (ocr *OCR) BankStatementRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "bank-statement", input)
}

// SKPR performs OCR on the given file using SKPR model
func (ocr *OCR) SKPR(ctx context.Context, input glair.OCRInput) (SKPR, error) {
	return recognize[SKPR](ctx, ocr, "skpr", input)
}

// SKPRRaw performs OCR on the given file using SKPR model
// and returns the raw API response as bytes
func (ocr *OCR) SKPRRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ctx, ocr, "skpr", input)
}

// BPKB performs OCR on the given file using BPKB model
func (ocr *OCR) BPKB(ctx context.Context, input glair.BPKBInput) (BPKB, error) {
	file, err := internal.ReadFile(input.Image)
	if err != nil {
		return BPKB{}, err
	}

	params := internal.RequestParameters{
		Url:       ocr.config.GetEndpointURL("ocr", "bpkb"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": file,
			"page":  input.Page,
		},
	}

	return internal.MakeMultipartRequest[BPKB](ctx, params, ocr.config)
}

// BPKBRaw performs OCR on the given file using BPKB model
// and returns the raw API response as bytes
func (ocr *OCR) BPKBRaw(ctx context.Context, input glair.BPKBInput) ([]byte, error) {
	file, err := internal.ReadFile(input.Image)
	if err != nil {
		return nil, err
	}

	params := internal.RequestParameters{
		Url:       ocr.config.GetEndpointURL("ocr", "bpkb"),
		RequestID: input.RequestID,
		Body: map[string]any{
			"image": file,
			"page":  input.Page,
		},
	}

	return internal.MakeMultipartRequestRaw(ctx, params, ocr.config)
}

// KTPSessions sends session request for KTP OCR using the prebuilt web page
func (ocr *OCR) KTPSessions(ctx context.Context, input glair.SessionsInput) (glair.Session, error) {
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

	return internal.MakeJSONRequest[glair.Session](ctx, params, ocr.config)
}

// NPWPSessions sends session request for NPWP OCR using the prebuilt web page
func (ocr *OCR) NPWPSessions(ctx context.Context, input glair.SessionsInput) (glair.Session, error) {
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

	return internal.MakeJSONRequest[glair.Session](ctx, params, ocr.config)
}
