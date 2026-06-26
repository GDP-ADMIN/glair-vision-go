// Package ocr is a collection of functions and objects that interacts
// with GLAIR Vision OCR products and its results
package ocr

import (
	"context"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/internal"
)

type OCR struct {
	config *glair.Config
}

type OCRResult[T any] struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Read   T      `json:"read"`
}

type PhotoField struct {
	Confidence     *float32 `json:"confidence,omitempty"`
	Polygon        [][]int  `json:"polygon,omitempty"`
	PageIndex      *int     `json:"page_index,omitempty"`
	Value          *string  `json:"value,omitempty"`
}

type OCRField[T any] struct {
	Confidence     *float32 `json:"confidence,omitempty"`
	ConfidenceText *float32 `json:"confidence_text,omitempty"`
	Polygon        [][]int  `json:"polygon,omitempty"`
	PageIndex      *int     `json:"page_index,omitempty"`
	ValueOriginal  *string  `json:"value_original,omitempty"`
	Value          *T       `json:"value,omitempty"`
}

type OCRPageField[T any] struct {
	OCRField[T]
	PageIndex int64 `json:"page_index,omitempty"`
}

type QualityBooleanField struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Value      *bool    `json:"value,omitempty"`
}

type BlurField struct {
	QualityBooleanField
	ConfidenceBlur *float64 `json:"confidence_blur,omitempty"`
}

type BrightField struct {
	QualityBooleanField
	ConfidenceBright *float64 `json:"confidence_bright,omitempty"`
}

type CropField struct {
	QualityBooleanField
	ConfidenceCrop *float64 `json:"confidence_crop,omitempty"`
}

type DarkField struct {
	QualityBooleanField
	ConfidenceDark *float64 `json:"confidence_dark,omitempty"`
}

type FlashField struct {
	QualityBooleanField
	ConfidenceFlash *float64 `json:"confidence_flash,omitempty"`
}

type PhotocopyField struct {
	QualityBooleanField
	ConfidencePhotocopy *float64 `json:"confidence_photocopy,omitempty"`
}

type ScreenField struct {
	QualityBooleanField
	ConfidenceScreen *float64 `json:"confidence_screen,omitempty"`
}

type DocumentField struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Value      *string  `json:"value,omitempty"`
}

type RotateField struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Value      *string  `json:"value,omitempty"`
	FixedImage *string  `json:"fixed_image,omitempty"`
}

type OCRQualities struct {
	Read   OCRQualitiesData  `json:"read,omitempty"`
	Status string            `json:"status,omitempty"`
	Reason string            `json:"reason,omitempty"`
}

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

func New(config *glair.Config) *OCR {
	return &OCR{
		config: config,
	}
}

func recognize[T any](ocr *OCR, ctx context.Context, endpoint string, input glair.OCRInput) (T, error) {
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

func recognizeRaw(ocr *OCR, ctx context.Context, endpoint string, input glair.OCRInput) ([]byte, error) {
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

func (ocr *OCR) KTP(ctx context.Context, input glair.OCRInput) (KTP, error) {
	return recognize[KTP](ocr, ctx, "ktp", input)
}

func (ocr *OCR) KTPRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "ktp", input)
}

func (ocr *OCR) KTPWithQuality(ctx context.Context, input glair.OCRInput) (KTPWithQuality, error) {
	return recognize[KTPWithQuality](ocr, ctx, "ktp/qualities", input)
}

func (ocr *OCR) KTPWithQualityRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "ktp/qualities", input)
}

func (ocr *OCR) NPWP(ctx context.Context, input glair.OCRInput) (NPWP, error) {
	return recognize[NPWP](ocr, ctx, "npwp", input)
}

func (ocr *OCR) NPWPRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "npwp", input)
}

func (ocr *OCR) KK(ctx context.Context, input glair.OCRInput) (KK, error) {
	return recognize[KK](ocr, ctx, "kk", input)
}

func (ocr *OCR) KKRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "kk", input)
}

func (ocr *OCR) STNK(ctx context.Context, input glair.OCRInput) (STNK, error) {
	return recognize[STNK](ocr, ctx, "stnk", input)
}

func (ocr *OCR) STNKRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "stnk", input)
}

func (ocr *OCR) SIM(ctx context.Context, input glair.OCRInput) (SIM, error) {
	return recognize[SIM](ocr, ctx, "sim", input)
}

func (ocr *OCR) SIMRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "sim", input)
}

func (ocr *OCR) Passport(ctx context.Context, input glair.OCRInput) (Passport, error) {
	return recognize[Passport](ocr, ctx, "passport", input)
}

func (ocr *OCR) PassportRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "passport", input)
}

func (ocr *OCR) Plate(ctx context.Context, input glair.OCRInput) (Plate, error) {
	return recognize[Plate](ocr, ctx, "plate", input)
}

func (ocr *OCR) PlateRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "plate", input)
}

func (ocr *OCR) GeneralDocument(ctx context.Context, input glair.OCRInput) (GeneralDocument, error) {
	return recognize[GeneralDocument](ocr, ctx, "general-document", input)
}

func (ocr *OCR) GeneralDocumentRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "general-document", input)
}

func (ocr *OCR) Invoice(ctx context.Context, input glair.OCRInput) (Invoice, error) {
	return recognize[Invoice](ocr, ctx, "invoice", input)
}

func (ocr *OCR) InvoiceRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "invoice", input)
}

func (ocr *OCR) Receipt(ctx context.Context, input glair.OCRInput) (Receipt, error) {
	return recognize[Receipt](ocr, ctx, "receipt", input)
}

func (ocr *OCR) ReceiptRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "receipt", input)
}

func (ocr *OCR) BankStatement(ctx context.Context, input glair.OCRInput) (BankStatement, error) {
	return recognize[BankStatement](ocr, ctx, "bank-statement", input)
}

func (ocr *OCR) BankStatementRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "bank-statement", input)
}

func (ocr *OCR) SKPR(ctx context.Context, input glair.OCRInput) (SKPR, error) {
	return recognize[SKPR](ocr, ctx, "skpr", input)
}

func (ocr *OCR) SKPRRaw(ctx context.Context, input glair.OCRInput) ([]byte, error) {
	return recognizeRaw(ocr, ctx, "skpr", input)
}

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
