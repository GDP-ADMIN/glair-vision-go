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
	BerlakuHingga    OCRField[string] `json:"berlaku_hingga,omitempty"`
	GolonganDarah    OCRField[string] `json:"golongan_darah,omitempty"`
	JenisKelamin     OCRField[string] `json:"jenis_kelamin,omitempty"`
	Kecamatan        OCRField[string] `json:"kecamatan,omitempty"`
	KelurahanDesa    OCRField[string] `json:"kelurahan_desa,omitempty"`
	Kewarganegaraan  OCRField[string] `json:"kewarganegaraan,omitempty"`
	KotaKabupaten    OCRField[string] `json:"kota_kabupaten,omitempty"`
	Nama             OCRField[string] `json:"nama,omitempty"`
	Nik              OCRField[string] `json:"nik,omitempty"`
	Pekerjaan        OCRField[string] `json:"pekerjaan,omitempty"`
	Provinsi         OCRField[string] `json:"provinsi,omitempty"`
	RtRw             OCRField[string] `json:"rt_rw,omitempty"`
	StatusPerkawinan OCRField[string] `json:"status_perkawinan,omitempty"`
	TanggalLahir     OCRField[string] `json:"tanggal_lahir,omitempty"`
	TempatLahir      OCRField[string] `json:"tempat_lahir,omitempty"`
	Foto 			 OCRField[string] `json:"foto,omitempty"`
	TandaTangan  	 OCRField[string] `json:"tanda_tangan,omitempty"`
}
