package ocr

type KK = OCRResult[KKData]

type KKData struct {
	Alamat             string  `json:"alamat,omitempty"`
	DesaKelurahan      string  `json:"desa_kelurahan,omitempty"`
	KabupatenKota      string  `json:"kabupaten_kota,omitempty"`
	Kecamatan          string  `json:"kecamatan,omitempty"`
	KodePos            string  `json:"kode_pos,omitempty"`
	NamaKepalaKeluarga string  `json:"nama_kepala_keluarga,omitempty"`
	NomorBlanko        string  `json:"nomor_blanko,omitempty"`
	NomorKK            string  `json:"nomor_kk,omitempty"`
	Provinsi           string  `json:"provinsi,omitempty"`
	RtRw               string  `json:"rt_rw,omitempty"`
	Table              KKTable `json:"table,omitempty"`
}

type KKTable struct {
	Agama                       string `json:"agama,omitempty"`
	GolonganDarah               string `json:"golongan_darah,omitempty"`
	JenisKelamin                string `json:"jenis_kelamin,omitempty"`
	Kewarganegaraan             string `json:"kewarganegaraan,omitempty"`
	NamaAyah                    string `json:"nama_ayah,omitempty"`
	NamaIbu                     string `json:"nama_ibu,omitempty"`
	NamaLengkap                 string `json:"nama_lengkap,omitempty"`
	Nik                         string `json:"nik,omitempty"`
	No                          string `json:"no,omitempty"`
	NoKitasKitap                string `json:"no_kitas_kitap,omitempty"`
	NoPaspor                    string `json:"no_paspor,omitempty"`
	Pendidikan                  string `json:"pendidikan,omitempty"`
	StatusHubunganDalamKeluarga string `json:"status_hubungan_dalam_keluarga,omitempty"`
	TanggalLahir                string `json:"tanggal_lahir,omitempty"`
	TanggalPerkawinan           string `json:"tanggal_perkawinan,omitempty"`
	TempatLahir                 string `json:"tempat_lahir,omitempty"`
}
