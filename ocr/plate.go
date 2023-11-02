package ocr

type Plate = OCRResult[PlatesData]

type PlatesData struct {
	Plates []string `json:"plates,omitempty"`
}
