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
	Alamat           OCRPageField[string] `json:"alamat,omitempty"`
	AlamatEmail      OCRPageField[string] `json:"alamat_email,omitempty"`
	Dikeluarkan      OCRPageField[string] `json:"dikeluarkan,omitempty"`
	NamaPemilik      OCRPageField[string] `json:"nama_pemilik,omitempty"`
	NomorBpkb        OCRPageField[string] `json:"nomor_bpkb,omitempty"`
	NoKtpTdp         OCRPageField[string] `json:"nomor_ktp_tdp,omitempty"`
	NoTelepon        OCRPageField[string] `json:"no_telepon,omitempty"`
	PadaTanggal      OCRPageField[string] `json:"pada_tanggal,omitempty"`
	Pekerjaan        OCRPageField[string] `json:"pekerjaan,omitempty"`
	StempelNomorBpkb OCRPageField[string] `json:"stempel_nomor_bpkb,omitempty"`
}

type VehicleIdentity struct {
	BahanBakar      OCRPageField[string] `json:"bahan_bakar,omitempty"`
	IsiSilinder     OCRPageField[string] `json:"isi_silinder,omitempty"`
	Jenis           OCRPageField[string] `json:"jenis,omitempty"`
	JumlahRoda      OCRPageField[string] `json:"jumlah_roda,omitempty"`
	JumlahSumbu     OCRPageField[string] `json:"jumlah_sumbu,omitempty"`
	Merk            OCRPageField[string] `json:"merk,omitempty"`
	Model           OCRPageField[string] `json:"model,omitempty"`
	NomorMesin      OCRPageField[string] `json:"nomor_mesin,omitempty"`
	NomorRangka     OCRPageField[string] `json:"nomor_rangka,omitempty"`
	NomorRegistrasi OCRPageField[string] `json:"nomor_registrasi,omitempty"`
	TahunPembuatan  OCRPageField[string] `json:"tahun_pembuatan,omitempty"`
	Type            OCRPageField[string] `json:"type,omitempty"`
	Warna           OCRPageField[string] `json:"warna,omitempty"`
	WarnaTnkb       OCRPageField[string] `json:"warna_tnkb,omitempty"`
}

type RegistrationDocument struct {
	NamaApm       OCRPageField[string] `json:"nama_apm,omitempty"`
	NomorFaktur   OCRPageField[string] `json:"nomor_faktur,omitempty"`
	NomorFormAbc  OCRPageField[string] `json:"nomor_form_abc,omitempty"`
	TanggalFaktur OCRPageField[string] `json:"tanggal_faktur,omitempty"`
}

type LastPageInformation struct {
	DiterbitkanOleh OCRPageField[string] `json:"diterbitkan_oleh,omitempty"`
	NoRegister      OCRPageField[string] `json:"no_register,omitempty"`
}
