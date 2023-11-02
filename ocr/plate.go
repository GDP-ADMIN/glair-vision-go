package ocr

// Plate stores OCR result of Plate model from the given input
type Plate = OCRResult[PlatesData]

type PlatesData struct {
	Plates []string `json:"plates,omitempty"`
}
