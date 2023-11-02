package ocr

type NPWP = OCRResult[NPWPData]

type NPWPData struct {
	Alamat string `json:"alamat,omitempty"`
	Nama   string `json:"nama,omitempty"`
	Nik    string `json:"nik,omitempty"`
	NoNPWP string `json:"noNpwp,omitempty"`
}
