package ocr

// SIM stores OCR result of SIM model from the given input
type SIM struct {
	OCRResult[SIMData]
	Images OCRImage `json:"images,omitempty"`
}

type SIMData struct {
	Alamat        OCRStringField `json:"alamat,omitempty"`
	Berlaku       OCRStringField `json:"berlaku,omitempty"`
	GolonganDarah OCRStringField `json:"golongan_darah,omitempty"`
	JenisKelamin  OCRStringField `json:"jenis_kelamin,omitempty"`
	Nama          OCRStringField `json:"nama,omitempty"`
	NomorSim      OCRStringField `json:"nomor_sim,omitempty"`
	Pekerjaan     OCRStringField `json:"pekerjaan,omitempty"`
	TanggalLahir  OCRStringField `json:"tangga_lahir,omitempty"`
	TempatLahir   OCRStringField `json:"tempat_lahir,omitempty"`
	Tinggi        OCRStringField `json:"tinggi,omitempty"`
	TipeSim       OCRStringField `json:"tipe_sim,omitempty"`
	Wilayah       OCRStringField `json:"wilayah,omitempty"`
}
