package ocr

// NPWP stores OCR result of NPWP model from the given input
type NPWP = OCRResult[NPWPData]

type NPWPData struct {
	Alamat OCRField `json:"alamat,omitempty"`
	Nama   OCRField `json:"nama,omitempty"`
	Nik    OCRField `json:"nik,omitempty"`
	NoNPWP OCRField `json:"noNpwp,omitempty"`
}
