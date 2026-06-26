package ocr

// STNK stores OCR result of STNK model from the given input.
type STNK = OCRResult[STNKData]

// STNKWithTable stores OCR result of STNK model with tax payment table.
type STNKWithTable = OCRResult[STNKWithTableData]

type STNKData struct {
	HalamanPajakTahunan HalamanPajakTahunanData  `json:"halaman_pajak_tahunan,omitempty"`
	HalamanSTNK         HalamanSTNKData `json:"halaman_stnk,omitempty"`
}

type STNKWithTableData struct {
	HalamanPajakTahunan HalamanPajakTahunanTable `json:"halaman_pajak_tahunan,omitempty"`
	HalamanSTNK         HalamanSTNKData         `json:"halaman_stnk,omitempty"`
}

type HalamanPajakTahunanData struct {
	NomorPajakTahunan     OCRPageField[string] `json:"nomor_pajak_tahunan,omitempty"`
	SamsatProvinsi        OCRPageField[string] `json:"samsat_provinsi,omitempty"`
	NomorRegistrasi       OCRPageField[string] `json:"nomor_registrasi,omitempty"`
	NamaPemilik           OCRPageField[string] `json:"nama_pemilik,omitempty"`
	Alamat                OCRPageField[string] `json:"alamat,omitempty"`
	Merk                  OCRPageField[string] `json:"merk,omitempty"`
	Tipe                  OCRPageField[string] `json:"tipe,omitempty"`
	IsiSilinder           OCRPageField[string] `json:"isi_silinder,omitempty"`
	Jenis                 OCRPageField[string] `json:"jenis,omitempty"`
	BahanBakar            OCRPageField[string] `json:"bahan_bakar,omitempty"`
	Model                 OCRPageField[string] `json:"model,omitempty"`
	WarnaTNKB             OCRPageField[string] `json:"warna_tnkb,omitempty"`
	TahunPembuatan        OCRPageField[string] `json:"tahun_pembuatan,omitempty"`
	TahunRegistrasi       OCRPageField[string] `json:"tahun_registrasi,omitempty"`
	Warna                 OCRPageField[string] `json:"warna,omitempty"`
	NomorBPKB             OCRPageField[string] `json:"nomor_bpkb,omitempty"`
	NomorRangka           OCRPageField[string] `json:"nomor_rangka,omitempty"`
	Ident                 OCRPageField[string] `json:"ident,omitempty"`
	NomorMesin            OCRPageField[string] `json:"nomor_mesin,omitempty"`
	BerlakuSampai         OCRPageField[string] `json:"berlaku_sampai,omitempty"`
	DitetapkanTanggal     OCRPageField[string] `json:"ditetapkan_tanggal,omitempty"`
	JumlahSumbu           OCRPageField[string] `json:"jumlah_sumbu,omitempty"`
	KepemilikanKe         OCRPageField[string] `json:"kepemilikan_ke,omitempty"`
	KodeNJKB              OCRPageField[string] `json:"kode_njkb,omitempty"`
	NomorIdentitasPemilik OCRPageField[string] `json:"nomor_identitas_pemilik,omitempty"`
	NomorRegistrasiLama   OCRPageField[string] `json:"nomor_registrasi_lama,omitempty"`
}

type HalamanSTNKData struct {
	NomorSTNK             OCRPageField[string] `json:"nomor_stnk,omitempty"`
	PoldaPenerbitan       OCRPageField[string] `json:"polda_penerbitan,omitempty"`
	NomorRegistrasi       OCRPageField[string] `json:"nomor_registrasi,omitempty"`
	NamaPemilik           OCRPageField[string] `json:"nama_pemilik,omitempty"`
	Alamat                OCRPageField[string] `json:"alamat,omitempty"`
	Merk                  OCRPageField[string] `json:"merk,omitempty"`
	Warna                 OCRPageField[string] `json:"warna,omitempty"`
	Tipe                  OCRPageField[string] `json:"tipe,omitempty"`
	BahanBakar            OCRPageField[string] `json:"bahan_bakar,omitempty"`
	Jenis                 OCRPageField[string] `json:"jenis,omitempty"`
	WarnaTNKB             OCRPageField[string] `json:"warna_tnkb,omitempty"`
	Model                 OCRPageField[string] `json:"model,omitempty"`
	TahunRegistrasi       OCRPageField[string] `json:"tahun_registrasi,omitempty"`
	TahunPembuatan        OCRPageField[string] `json:"tahun_pembuatan,omitempty"`
	NomorBPKB             OCRPageField[string] `json:"nomor_bpkb,omitempty"`
	IsiSilinder           OCRPageField[string] `json:"isi_silinder,omitempty"`
	KodeLokasi            OCRPageField[string] `json:"kode_lokasi,omitempty"`
	NomorUrutPendaftaran  OCRPageField[string] `json:"nomor_urut_pendaftaran,omitempty"`
	TempatPenerbitan      OCRPageField[string] `json:"tempat_penerbitan,omitempty"`
	TanggalPenerbitan     OCRPageField[string] `json:"tanggal_penerbitan,omitempty"`
	BerlakuSampai         OCRPageField[string] `json:"berlaku_sampai,omitempty"`
	NomorIdentitasPemilik OCRPageField[string] `json:"nomor_identitas_pemilik,omitempty"`
	NomorMesin            OCRPageField[string] `json:"nomor_mesin,omitempty"`
	NomorRangka           OCRPageField[string] `json:"nomor_rangka,omitempty"`
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
