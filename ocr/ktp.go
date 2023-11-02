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
	Agama            OCRField `json:"agama,omitempty"`
	Alamat           OCRField `json:"alamat,omitempty"`
	BerlakuHingga    OCRField `json:"berlakuHingga,omitempty"`
	GolonganDarah    OCRField `json:"golonganDarah,omitempty"`
	JenisKelamin     OCRField `json:"jenisKelamin,omitempty"`
	Kecamatan        OCRField `json:"kecamatan,omitempty"`
	KelurahanDesa    OCRField `json:"kelurahanDesa,omitempty"`
	Kewarganegaraan  OCRField `json:"kewarganegaraan,omitempty"`
	KotaKabupaten    OCRField `json:"kotaKabupaten,omitempty"`
	Nama             OCRField `json:"nama,omitempty"`
	Nik              OCRField `json:"nik,omitempty"`
	Pekerjaan        OCRField `json:"pekerjaan,omitempty"`
	Provinsi         OCRField `json:"provinsi,omitempty"`
	RtRw             OCRField `json:"rtRw,omitempty"`
	StatusPerkawinan OCRField `json:"statusPerkawinan,omitempty"`
	TanggalLahir     OCRField `json:"tanggalLahir,omitempty"`
	TempatLahir      OCRField `json:"tempatLahir,omitempty"`
}
