package ocr

// NPWP stores OCR result of NPWP model from the given input
type NPWP = OCRResult[NPWPData]

type NPWPData struct {
	Alamat OCRStringField `json:"alamat,omitempty"`
	Nama   OCRStringField `json:"nama,omitempty"`
	Nik    OCRStringField `json:"nik,omitempty"`
	NoNPWP OCRStringField `json:"noNpwp,omitempty"`
}
