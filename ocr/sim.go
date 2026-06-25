package ocr

// SIM stores OCR result of SIM model from the given input
type SIM struct {
	OCRResult[SIMData]
	Images OCRImage `json:"images,omitempty"`
}

type SIMData struct {
	Alamat        OCRField[string] `json:"alamat,omitempty"`
	Berlaku       OCRField[string] `json:"berlaku,omitempty"`
	GolonganDarah OCRField[string] `json:"golongan_darah,omitempty"`
	JenisKelamin  OCRField[string] `json:"jenis_kelamin,omitempty"`
	Nama          OCRField[string] `json:"nama,omitempty"`
	NomorSim      OCRField[string] `json:"nomor_sim,omitempty"`
	Pekerjaan     OCRField[string] `json:"pekerjaan,omitempty"`
	TanggalLahir  OCRField[string] `json:"tanggal_lahir,omitempty"`
	TempatLahir   OCRField[string] `json:"tempat_lahir,omitempty"`
	Tinggi        OCRField[string] `json:"tinggi,omitempty"`
	TipeSim       OCRField[string] `json:"tipe_sim,omitempty"`
	Wilayah       OCRField[string] `json:"wilayah,omitempty"`
	Foto          OCRField[string] `json:"foto,omitempty"`
	TandaTangan   OCRField[string] `json:"tanda_tangan,omitempty"`
}
