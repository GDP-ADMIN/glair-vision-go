package ocr

// STNK stores OCR result of STNK model from the given input.
type STNK = OCRResult[STNKData]

// STNKWithTable stores OCR result of STNK model with tax payment table.
type STNKWithTable = OCRResult[STNKWithTableData]

type STNKData struct {
	HalamanPajakTahunan HalamanPajakTahunanData `json:"halaman_pajak_tahunan,omitempty"`
	HalamanSTNK         HalamanSTNKData         `json:"halaman_stnk,omitempty"`
}

type STNKWithTableData struct {
	HalamanPajakTahunan HalamanPajakTahunanTable `json:"halaman_pajak_tahunan,omitempty"`
	HalamanSTNK         HalamanSTNKData          `json:"halaman_stnk,omitempty"`
}

type HalamanPajakTahunanData struct {
	NomorPajakTahunan     OCRField[string] `json:"nomor_pajak_tahunan,omitempty"`
	SamsatProvinsi        OCRField[string] `json:"samsat_provinsi,omitempty"`
	NomorRegistrasi       OCRField[string] `json:"nomor_registrasi,omitempty"`
	NamaPemilik           OCRField[string] `json:"nama_pemilik,omitempty"`
	Alamat                OCRField[string] `json:"alamat,omitempty"`
	Merk                  OCRField[string] `json:"merk,omitempty"`
	Tipe                  OCRField[string] `json:"tipe,omitempty"`
	IsiSilinder           OCRField[string] `json:"isi_silinder,omitempty"`
	Jenis                 OCRField[string] `json:"jenis,omitempty"`
	BahanBakar            OCRField[string] `json:"bahan_bakar,omitempty"`
	Model                 OCRField[string] `json:"model,omitempty"`
	WarnaTNKB             OCRField[string] `json:"warna_tnkb,omitempty"`
	TahunPembuatan        OCRField[string] `json:"tahun_pembuatan,omitempty"`
	TahunRegistrasi       OCRField[string] `json:"tahun_registrasi,omitempty"`
	Warna                 OCRField[string] `json:"warna,omitempty"`
	NomorBPKB             OCRField[string] `json:"nomor_bpkb,omitempty"`
	NomorRangka           OCRField[string] `json:"nomor_rangka,omitempty"`
	Ident                 OCRField[string] `json:"ident,omitempty"`
	NomorMesin            OCRField[string] `json:"nomor_mesin,omitempty"`
	BerlakuSampai         OCRField[string] `json:"berlaku_sampai,omitempty"`
	DitetapkanTanggal     OCRField[string] `json:"ditetapkan_tanggal,omitempty"`
	JumlahSumbu           OCRField[string] `json:"jumlah_sumbu,omitempty"`
	KepemilikanKe         OCRField[string] `json:"kepemilikan_ke,omitempty"`
	KodeNJKB              OCRField[string] `json:"kode_njkb,omitempty"`
	NomorIdentitasPemilik OCRField[string] `json:"nomor_identitas_pemilik,omitempty"`
	NomorRegistrasiLama   OCRField[string] `json:"nomor_registrasi_lama,omitempty"`
}

type HalamanSTNKData struct {
	NomorSTNK             OCRField[string] `json:"nomor_stnk,omitempty"`
	PoldaPenerbitan       OCRField[string] `json:"polda_penerbitan,omitempty"`
	NomorRegistrasi       OCRField[string] `json:"nomor_registrasi,omitempty"`
	NamaPemilik           OCRField[string] `json:"nama_pemilik,omitempty"`
	Alamat                OCRField[string] `json:"alamat,omitempty"`
	Merk                  OCRField[string] `json:"merk,omitempty"`
	Warna                 OCRField[string] `json:"warna,omitempty"`
	Tipe                  OCRField[string] `json:"tipe,omitempty"`
	BahanBakar            OCRField[string] `json:"bahan_bakar,omitempty"`
	Jenis                 OCRField[string] `json:"jenis,omitempty"`
	WarnaTNKB             OCRField[string] `json:"warna_tnkb,omitempty"`
	Model                 OCRField[string] `json:"model,omitempty"`
	TahunRegistrasi       OCRField[string] `json:"tahun_registrasi,omitempty"`
	TahunPembuatan        OCRField[string] `json:"tahun_pembuatan,omitempty"`
	NomorBPKB             OCRField[string] `json:"nomor_bpkb,omitempty"`
	IsiSilinder           OCRField[string] `json:"isi_silinder,omitempty"`
	KodeLokasi            OCRField[string] `json:"kode_lokasi,omitempty"`
	NomorUrutPendaftaran  OCRField[string] `json:"nomor_urut_pendaftaran,omitempty"`
	TempatPenerbitan      OCRField[string] `json:"tempat_penerbitan,omitempty"`
	TanggalPenerbitan     OCRField[string] `json:"tanggal_penerbitan,omitempty"`
	BerlakuSampai         OCRField[string] `json:"berlaku_sampai,omitempty"`
	NomorIdentitasPemilik OCRField[string] `json:"nomor_identitas_pemilik,omitempty"`
	NomorMesin            OCRField[string] `json:"nomor_mesin,omitempty"`
	NomorRangka           OCRField[string] `json:"nomor_rangka,omitempty"`
}

type HalamanPajakTahunanTable struct {
	HalamanPajakTahunanData
	Table []TableData `json:"table,omitempty"`
}

type TableData struct {
	JenisBiaya          OCRField[string] `json:"jenis_biaya,omitempty"`
	Pokok               OCRField[string] `json:"pokok,omitempty"`
	SanksiAdministratif OCRField[string] `json:"sanksi_administratif,omitempty"`
	Jumlah              OCRField[string] `json:"jumlah,omitempty"`
}
