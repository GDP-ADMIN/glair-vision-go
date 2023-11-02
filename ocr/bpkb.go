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
	Alamat      string `json:"alamat,omitempty"`
	AlamatEmail string `json:"alamat_email,omitempty"`
	Dikeluarkan string `json:"dikeluarkan,omitempty"`
	NamaPemilik string `json:"nama_pemilik,omitempty"`
	NomorBpkb   string `json:"nomor_bpkb,omitempty"`
	NoKtpTdp    string `json:"nomor_ktp_tdp,omitempty"`
	NoTelepon   string `json:"no_telepon,omitempty"`
	PadaTanggal string `json:"pada_tanggal,omitempty"`
	Pekerjaan   string `json:"pekerjaan,omitempty"`
}

type VehicleIdentity struct {
	BahanBakar      string `json:"bahan_bakar,omitempty"`
	IsiSilinder     string `json:"isi_silinder,omitempty"`
	Jenis           string `json:"jenis,omitempty"`
	JumlahRoda      string `json:"jumlah_roda,omitempty"`
	JumlahSumbu     string `json:"jumlah_sumbu,omitempty"`
	Merk            string `json:"merk,omitempty"`
	Model           string `json:"model,omitempty"`
	NomorMesin      string `json:"nomor_mesin,omitempty"`
	NomorRangka     string `json:"nomor_rangka,omitempty"`
	NomorRegistrasi string `json:"nomor_registrasi,omitempty"`
	TahunPembuatan  string `json:"tahun_pembuatan,omitempty"`
	Type            string `json:"type,omitempty"`
	Warna           string `json:"warna,omitempty"`
	WarnaTnkb       string `json:"warna_tnkb,omitempty"`
}

type RegistrationDocument struct {
	NamaApm       string `json:"nama_apm,omitempty"`
	NomorFaktur   string `json:"nomor_faktur,omitempty"`
	NomorFormAbc  string `json:"nomor_form_abc,omitempty"`
	TanggalFaktur string `json:"tanggal_faktur,omitempty"`
}

type LastPageInformation struct {
	DiterbitkanOleh string `json:"diterbitkan_oleh,omitempty"`
	NoRegister      string `json:"no_register,omitempty"`
}
