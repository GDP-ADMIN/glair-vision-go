package ocr

// NPWP stores OCR result of NPWP model from the given input
type NPWP = OCRResult[NPWPData]

type NPWPData struct {
	Alamat OCRField[string] `json:"alamat,omitempty"`
	Nama   OCRField[string] `json:"nama,omitempty"`
	Nik    OCRField[string] `json:"nik,omitempty"`
	NoNPWP OCRField[string] `json:"no_npwp,omitempty"`
}
