package ocr

type GeneralDocument = OCRResult[GeneralDocumentData]

type GeneralDocumentData struct {
	AllTexts OCRField `json:"all_texts,omitempty"`
}
