package model

// CoWorkingSpace menyimpan data tempat kerja (Aturan: Menggunakan Struct/Record)
type CoWorkingSpace struct {
	ID        int
	Nama      string
	Lokasi    string
	HargaSewa float64 // Harga sewa per jam
	Rating    float64 // Nilai rata-rata ulasan dari skala 1-5
	HasWiFi   bool    // Fasilitas penanda wifi
}

// Review menyimpan ulasan langsung dari pengunjung
type Review struct {
	ID              int
	SpaceID         int // Menghubungkan review ke CoWorkingSpace lewat ID
	NamaPengguna    string
	Komentar        string
	RatingDiberikan float64
}
