package ocr

import (
	"os"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/internal"
)

// KTP is an object that stores OCR result from the provided KTP
type KTP struct {
	OCRResult[KTPRead]
	Image OCRImage `json:"image,omitempty"`
}

type KTPWithQuality struct {
	KTP
	Qualities OCRQualities `json:"qualities,omitempty"`
}

type KTPRead struct {
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

// Ktp
func (ocr *OCR) Ktp(file *os.File) (KTP, error) {
	url, err := ocr.config.GetEndpointURL("ocr", "ktp")
	if err != nil {
		return KTP{}, glair.ErrInvalidBaseUrl
	}

	if file == nil {
		return KTP{}, glair.ErrFileRequired
	}

	return internal.MakeRequest[KTP](ocr.config, url, file)
}

// KtpWithQuality
func (ocr *OCR) KtpWithQuality(file *os.File) (KTPWithQuality, error) {
	url, err := ocr.config.GetEndpointURL("ocr", "ktp")
	if err != nil {
		return KTPWithQuality{}, glair.ErrInvalidBaseUrl
	}

	if file == nil {
		return KTPWithQuality{}, glair.ErrFileRequired
	}

	return internal.MakeRequest[KTPWithQuality](ocr.config, url, file)
}
