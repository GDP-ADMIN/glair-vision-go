package ocr

import (
	"os"

	"github.com/glair-ai/glair-vision-go/internal"
)

// KTP is an object that stores OCR result from the provided KTP
type KTP struct {
	OCRResult[KTPRead]
	Image OCRImage
}

type KTPWithQuality struct {
	KTP
	Qualities OCRQualities
}

type KTPRead struct {
	Agama            OCRField `json:"agama"`
	Alamat           OCRField `json:"alamat"`
	BerlakuHingga    OCRField `json:"berlakuHingga"`
	GolonganDarah    OCRField `json:"golonganDarah"`
	JenisKelamin     OCRField `json:"jenisKelamin"`
	Kecamatan        OCRField `json:"kecamatan"`
	KelurahanDesa    OCRField `json:"kelurahanDesa"`
	Kewarganegaraan  OCRField `json:"kewarganegaraan"`
	KotaKabupaten    OCRField `json:"kotaKabupaten"`
	Nama             OCRField `json:"nama"`
	Nik              OCRField `json:"nik"`
	Pekerjaan        OCRField `json:"pekerjaan"`
	Provinsi         OCRField `json:"provinsi"`
	RtRw             OCRField `json:"rtRw"`
	StatusPerkawinan OCRField `json:"statusPerkawinan"`
	TanggalLahir     OCRField `json:"tanggalLahir"`
	TempatLahir      OCRField `json:"tempatLahir"`
}

// Ktp calls
func (ocr *OCR) Ktp(file *os.File) (KTP, error) {
	return internal.MakeRequest[KTP](ocr.config, "ktp", file)
}

// KtpWithQuality
func (ocr *OCR) KtpWithQuality(file *os.File) (KTPWithQuality, error) {
	return internal.MakeRequest[KTPWithQuality](ocr.config, "ktp/qualities", file)
}
