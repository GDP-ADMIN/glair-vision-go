package ocr

type BPKB = OCRResult[BPKBRead]

type BPKBRead struct {
	IdentitasPemilik IdentitasPemilik `json:"identitas_pemilik,omitempty"`
}

type IdentitasPemilik struct {
	Alamat string `json:"alamat,omitempty"`
}
