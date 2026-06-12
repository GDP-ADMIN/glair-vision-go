package ocr

// STNK stores OCR result of STNK model from the given input
type STNK = OCRResult[STNKData]

type STNKData struct {
	Alamat               OCRField[string] `json:"alamat,omitempty"`
	BahanBakar           OCRField[string] `json:"bahan_bakar,omitempty"`
	BerlakuSampai        OCRField[string] `json:"berlaku_sampai,omitempty"`
	IsiSilinder          OCRField[string] `json:"isi_silinder,omitempty"`
	Jenis                OCRField[string] `json:"jenis,omitempty"`
	KodeLokasi           OCRField[string] `json:"kode_lokasi,omitempty"`
	Merk                 OCRField[string] `json:"merk,omitempty"`
	Model                OCRField[string] `json:"model,omitempty"`
	NamaPemilik          OCRField[string] `json:"nama_pemilik,omitempty"`
	NomorBpkb            OCRField[string] `json:"nomor_bpkb,omitempty"`
	NomorMesin           OCRField[string] `json:"nomor_mesin,omitempty"`
	NomorRegistrasi      OCRField[string] `json:"nomor_registrasi,omitempty"`
	NomorStnk            OCRField[string] `json:"nomor_stnk,omitempty"`
	NomorUrutPendaftaran OCRField[string] `json:"nomor_urut_pendaftaran,omitempty"`
	TahunPembuatan       OCRField[string] `json:"tahun_pembuatan,omitempty"`
	TahunRegistrasi      OCRField[string] `json:"tahun_registrasi,omitempty"`
	Tipe                 OCRField[string] `json:"tipe,omitempty"`
	Warna                OCRField[string] `json:"warna,omitempty"`
	WarnaTnkb            OCRField[string] `json:"warna_tnkb,omitempty"`
}
