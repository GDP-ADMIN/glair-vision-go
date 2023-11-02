package ocr

// STNK stores OCR result of STNK model from the given input
type STNK = OCRResult[STNKData]

type STNKData struct {
	Alamat               string `json:"alamat,omitempty"`
	BahanBakar           string `json:"bahan_bakar,omitempty"`
	BerlakuSampai        string `json:"berlaku_sampai,omitempty"`
	IsiSilinder          string `json:"isi_silinder,omitempty"`
	Jenis                string `json:"jenis,omitempty"`
	KodeLokasi           string `json:"kode_lokasi,omitempty"`
	Merk                 string `json:"merk,omitempty"`
	Model                string `json:"model,omitempty"`
	NamaPemilik          string `json:"nama_pemilik,omitempty"`
	NomorBpkb            string `json:"nomor_bpkb,omitempty"`
	NomorMesin           string `json:"nomor_mesin,omitempty"`
	NomorRegistrasi      string `json:"nomor_registrasi,omitempty"`
	NomorStnk            string `json:"nomor_stnk,omitempty"`
	NomorUrutPendaftaran string `json:"nomor_urut_pendaftaran,omitempty"`
	TahunPembuatan       string `json:"tahun_pembuatan,omitempty"`
	TahunRegistrasi      string `json:"tahun_registrasi,omitempty"`
	Tipe                 string `json:"tipe,omitempty"`
	Warna                string `json:"warna,omitempty"`
	WarnaTnkb            string `json:"warna_tnkb"`
}
