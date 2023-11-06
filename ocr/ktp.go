package ocr

// KTP stores OCR result of KTP model from the given input
type KTP struct {
	OCRResult[KTPData]
	Image OCRImage `json:"image,omitempty"`
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
	Agama            OCRStringField `json:"agama,omitempty"`
	Alamat           OCRStringField `json:"alamat,omitempty"`
	BerlakuHingga    OCRStringField `json:"berlakuHingga,omitempty"`
	GolonganDarah    OCRStringField `json:"golonganDarah,omitempty"`
	JenisKelamin     OCRStringField `json:"jenisKelamin,omitempty"`
	Kecamatan        OCRStringField `json:"kecamatan,omitempty"`
	KelurahanDesa    OCRStringField `json:"kelurahanDesa,omitempty"`
	Kewarganegaraan  OCRStringField `json:"kewarganegaraan,omitempty"`
	KotaKabupaten    OCRStringField `json:"kotaKabupaten,omitempty"`
	Nama             OCRStringField `json:"nama,omitempty"`
	Nik              OCRStringField `json:"nik,omitempty"`
	Pekerjaan        OCRStringField `json:"pekerjaan,omitempty"`
	Provinsi         OCRStringField `json:"provinsi,omitempty"`
	RtRw             OCRStringField `json:"rtRw,omitempty"`
	StatusPerkawinan OCRStringField `json:"statusPerkawinan,omitempty"`
	TanggalLahir     OCRStringField `json:"tanggalLahir,omitempty"`
	TempatLahir      OCRStringField `json:"tempatLahir,omitempty"`
}
