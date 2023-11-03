package ocr

// Plate stores OCR result of Plate model from the given input
type Plate = OCRResult[PlatesData]

type PlatesData struct {
	Plates []struct {
		Text string `json:"text,omitempty"`
	} `json:"plates,omitempty"`
}
