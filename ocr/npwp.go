package ocr

type NPWP = OCRResult[NPWPRead]

type NPWPRead struct {
	Alamat string `json:"alamat,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Nik    string `json:"nik,omitempty"`
	NoNPWP string `json:"noNpwp,omitempty"`
}
