package ocr

// KTP stores OCR result of KTP model from the given input
type KTP struct {
	OCRResult[KTPData]
	Images OCRImage `json:"images,omitempty"`
}

// KTPWithQuality stores OCR result of KTP model from the given
// input with quality remarks
type KTPWithQuality struct {
	KTP
	Qualities OCRQualities `json:"qualities,omitempty"`
}

// KTPData stores key-value field of KTP that can be
// recognized by the OCR engine
type KTPData struct {
	Agama            OCRField[string] `json:"agama,omitempty"`
	Alamat           OCRField[string] `json:"alamat,omitempty"`
	BerlakuHingga    OCRField[string] `json:"berlakuHingga,omitempty"`
	GolonganDarah    OCRField[string] `json:"golonganDarah,omitempty"`
	JenisKelamin     OCRField[string] `json:"jenisKelamin,omitempty"`
	Kecamatan        OCRField[string] `json:"kecamatan,omitempty"`
	KelurahanDesa    OCRField[string] `json:"kelurahanDesa,omitempty"`
	Kewarganegaraan  OCRField[string] `json:"kewarganegaraan,omitempty"`
	KotaKabupaten    OCRField[string] `json:"kotaKabupaten,omitempty"`
	Nama             OCRField[string] `json:"nama,omitempty"`
	Nik              OCRField[string] `json:"nik,omitempty"`
	Pekerjaan        OCRField[string] `json:"pekerjaan,omitempty"`
	Provinsi         OCRField[string] `json:"provinsi,omitempty"`
	RtRw             OCRField[string] `json:"rtRw,omitempty"`
	StatusPerkawinan OCRField[string] `json:"statusPerkawinan,omitempty"`
	TanggalLahir     OCRField[string] `json:"tanggalLahir,omitempty"`
	TempatLahir      OCRField[string] `json:"tempatLahir,omitempty"`
}
