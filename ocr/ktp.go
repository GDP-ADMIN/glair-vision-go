package ocr

// KTP stores OCR result of KTP model from the given input
type KTP struct {
	OCRResult[KTPData]
}

// KTPWithQuality stores OCR result of KTP model from the given
// input with quality remarks
type KTPWithQuality struct {
	KTP
	Qualities OCRQualitiesData `json:"qualities,omitempty"`
}

// KTPData stores key-value field of KTP that can be
// recognized by the OCR engine
type KTPData struct {
	Agama            OCRPageField[string] `json:"agama,omitempty"`
	Alamat           OCRPageField[string] `json:"alamat,omitempty"`
	BerlakuHingga    OCRPageField[string] `json:"berlaku_hingga,omitempty"`
	GolonganDarah    OCRPageField[string] `json:"golongan_darah,omitempty"`
	JenisKelamin     OCRPageField[string] `json:"jenis_kelamin,omitempty"`
	Kecamatan        OCRPageField[string] `json:"kecamatan,omitempty"`
	KelurahanDesa    OCRPageField[string] `json:"kelurahan_desa,omitempty"`
	Kewarganegaraan  OCRPageField[string] `json:"kewarganegaraan,omitempty"`
	KotaKabupaten    OCRPageField[string] `json:"kota_kabupaten,omitempty"`
	Nama             OCRPageField[string] `json:"nama,omitempty"`
	Nik              OCRPageField[string] `json:"nik,omitempty"`
	Pekerjaan        OCRPageField[string] `json:"pekerjaan,omitempty"`
	Provinsi         OCRPageField[string] `json:"provinsi,omitempty"`
	RtRw             OCRPageField[string] `json:"rt_rw,omitempty"`
	StatusPerkawinan OCRPageField[string] `json:"status_perkawinan,omitempty"`
	TanggalLahir     OCRPageField[string] `json:"tanggal_lahir,omitempty"`
	TempatLahir      OCRPageField[string] `json:"tempat_lahir,omitempty"`
	Foto             PhotoField       `json:"foto,omitempty"`
	TandaTangan      PhotoField       `json:"tanda_tangan,omitempty"`
}
