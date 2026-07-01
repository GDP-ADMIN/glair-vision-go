package ocr

// KK stores OCR result of Kartu Keluarga model from the given input
type KK = OCRResult[KKData]

type KKData struct {
	Alamat             OCRField[string] `json:"alamat,omitempty"`
	DesaKelurahan      OCRField[string] `json:"desa_kelurahan,omitempty"`
	KabupatenKota      OCRField[string] `json:"kabupaten_kota,omitempty"`
	Kecamatan          OCRField[string] `json:"kecamatan,omitempty"`
	KodePos            OCRField[string] `json:"kode_pos,omitempty"`
	NamaKepalaKeluarga OCRField[string] `json:"nama_kepala_keluarga,omitempty"`
	NomorBlanko        OCRField[string] `json:"nomor_blanko,omitempty"`
	NomorKK            OCRField[string] `json:"nomor_kk,omitempty"`
	Provinsi           OCRField[string] `json:"provinsi,omitempty"`
	RtRw               OCRField[string] `json:"rt_rw,omitempty"`
	Table              []KKTable        `json:"table,omitempty"`
}

type KKTable struct {
	Agama                       OCRField[string] `json:"agama,omitempty"`
	GolonganDarah               OCRField[string] `json:"golongan_darah,omitempty"`
	JenisKelamin                OCRField[string] `json:"jenis_kelamin,omitempty"`
	Kewarganegaraan             OCRField[string] `json:"kewarganegaraan,omitempty"`
	NamaAyah                    OCRField[string] `json:"nama_ayah,omitempty"`
	NamaIbu                     OCRField[string] `json:"nama_ibu,omitempty"`
	NamaLengkap                 OCRField[string] `json:"nama_lengkap,omitempty"`
	Nik                         OCRField[string] `json:"nik,omitempty"`
	No                          OCRField[string] `json:"no,omitempty"`
	NoKitasKitap                OCRField[string] `json:"no_kitas_kitap,omitempty"`
	NoPaspor                    OCRField[string] `json:"no_paspor,omitempty"`
	Pendidikan                  OCRField[string] `json:"pendidikan,omitempty"`
	StatusHubunganDalamKeluarga OCRField[string] `json:"status_hubungan_dalam_keluarga,omitempty"`
	TanggalLahir                OCRField[string] `json:"tanggal_lahir,omitempty"`
	TanggalPerkawinan           OCRField[string] `json:"tanggal_perkawinan,omitempty"`
	TempatLahir                 OCRField[string] `json:"tempat_lahir,omitempty"`
}
