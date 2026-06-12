package ocr

// BPKB stores OCR result of BPKB model from the given input
type BPKB = OCRResult[BPKBData]

type BPKBData struct {
	IdentitasPemilik         OwnerIdentity        `json:"identitas_pemilik,omitempty"`
	IdentitasKendaraan       VehicleIdentity      `json:"identitas_kendaraan,omitempty"`
	DokumenRegistrasiPertama RegistrationDocument `json:"dokumen_registrasi_pertama,omitempty"`
	HalamanTerakhir          LastPageInformation  `json:"halaman_terakhir,omitempty"`
}

type OwnerIdentity struct {
	Alamat      OCRField[string] `json:"alamat,omitempty"`
	AlamatEmail OCRField[string] `json:"alamat_email,omitempty"`
	Dikeluarkan OCRField[string] `json:"dikeluarkan,omitempty"`
	NamaPemilik OCRField[string] `json:"nama_pemilik,omitempty"`
	NomorBpkb   OCRField[string] `json:"nomor_bpkb,omitempty"`
	NoKtpTdp    OCRField[string] `json:"nomor_ktp_tdp,omitempty"`
	NoTelepon   OCRField[string] `json:"no_telepon,omitempty"`
	PadaTanggal OCRField[string] `json:"pada_tanggal,omitempty"`
	Pekerjaan   OCRField[string] `json:"pekerjaan,omitempty"`
}

type VehicleIdentity struct {
	BahanBakar      OCRField[string] `json:"bahan_bakar,omitempty"`
	IsiSilinder     OCRField[string] `json:"isi_silinder,omitempty"`
	Jenis           OCRField[string] `json:"jenis,omitempty"`
	JumlahRoda      OCRField[string] `json:"jumlah_roda,omitempty"`
	JumlahSumbu     OCRField[string] `json:"jumlah_sumbu,omitempty"`
	Merk            OCRField[string] `json:"merk,omitempty"`
	Model           OCRField[string] `json:"model,omitempty"`
	NomorMesin      OCRField[string] `json:"nomor_mesin,omitempty"`
	NomorRangka     OCRField[string] `json:"nomor_rangka,omitempty"`
	NomorRegistrasi OCRField[string] `json:"nomor_registrasi,omitempty"`
	TahunPembuatan  OCRField[string] `json:"tahun_pembuatan,omitempty"`
	Type            OCRField[string] `json:"type,omitempty"`
	Warna           OCRField[string] `json:"warna,omitempty"`
	WarnaTnkb       OCRField[string] `json:"warna_tnkb,omitempty"`
}

type RegistrationDocument struct {
	NamaApm       OCRField[string] `json:"nama_apm,omitempty"`
	NomorFaktur   OCRField[string] `json:"nomor_faktur,omitempty"`
	NomorFormAbc  OCRField[string] `json:"nomor_form_abc,omitempty"`
	TanggalFaktur OCRField[string] `json:"tanggal_faktur,omitempty"`
}

type LastPageInformation struct {
	DiterbitkanOleh OCRField[string] `json:"diterbitkan_oleh,omitempty"`
	NoRegister      OCRField[string] `json:"no_register,omitempty"`
}
