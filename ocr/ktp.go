package ocr

import (
	"os"

	"github.com/glair-ai/glair-vision-go/internal"
)

type KTP struct {
	OCRResult[KTPRead]
	Image OCRImage
}

type KTPWithQuality struct {
	KTP
	Qualities OCRQualities
}

type KTPRead struct {
	Agama            OCRField
	Alamat           OCRField
	BerlakuHingga    OCRField
	GolonganDarah    OCRField
	JenisKelamin     OCRField
	Kecamatan        OCRField
	KelurahanDesa    OCRField
	Kewarganegaraan  OCRField
	KotaKabupaten    OCRField
	Nama             OCRField
	Nik              OCRField
	Pekerjaan        OCRField
	Provinsi         OCRField
	RtRw             OCRField
	StatusPerkawinan OCRField
	TanggalLahir     OCRField
	TempatLahir      OCRField
}

func (ocr *OCR) Ktp(file *os.File) (KTP, error) {
	return internal.MakeRequest[KTP](ocr.config, "ktp", file)
}

func (ocr *OCR) KtpWithQuality(file *os.File) {

}
