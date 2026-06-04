package main

import (
	"fmt"
	"tubes_coworking/alpro"
	"tubes_coworking/model"
)

var daftarSpace []model.CoWorkingSpace
var totalData int = 0

func main() {
	inisialisasiDataAwal()

	var pilihan int
	for {
		tampilkanMenu()
		fmt.Print("Pilih menu (1-7): ")
		fmt.Scanln(&pilihan)

		if pilihan == 7 {
			fmt.Println("\nAplikasi ditutup")
			break
		}

		prosesMenu(pilihan)
	}
}

func tampilkanMenu() {
	fmt.Println("\n=======================================")
	fmt.Println("   SISTEM INFORMASI CO-WORKING SPACE   ")
	fmt.Println("=======================================")
	fmt.Println("1. Tampilkan Semua Co-Working Space")
	fmt.Println("2. Tambah Data Co-Working Space (Admin)")
	fmt.Println("3. Ubah Data Co-Working Space (Admin)")
	fmt.Println("4. Hapus Data Co-Working Space (Admin)")
	fmt.Println("5. Urutkan Data Tempat (Sorting)")
	fmt.Println("6. Cari Data Tempat (Searching)")
	fmt.Println("7. Keluar dari Aplikasi")
	fmt.Println("=======================================")
}

func prosesMenu(pilihan int) {
	switch pilihan {
	case 1:
		tampilkanSemuaSpace()
	case 2:
		tambahSpaceBaru()
	case 3:
		ubahSpace()
	case 4:
		hapusSpace()
	case 5:
		menuSorting()
	case 6:
		menuSearching()
	default:
		fmt.Println("\nPilihan salah! Silakan masukkan angka 1 hingga 7.")
	}
}

func inisialisasiDataAwal() {
	daftarSpace = append(daftarSpace, model.CoWorkingSpace{
		ID: 1, Nama: "Cikini Creative Hub", Lokasi: "Jakarta", HargaSewa: 50000, Rating: 4.2, HasWiFi: true,
	})
	daftarSpace = append(daftarSpace, model.CoWorkingSpace{
		ID: 2, Nama: "Bogor Space & Co", Lokasi: "Bogor", HargaSewa: 35000, Rating: 4.8, HasWiFi: true,
	})
	daftarSpace = append(daftarSpace, model.CoWorkingSpace{
		ID: 3, Nama: "Kemang Workspace", Lokasi: "Jakarta", HargaSewa: 75000, Rating: 4.5, HasWiFi: false,
	})
	totalData = 3
}

func tampilkanSemuaSpace() {
	if len(daftarSpace) == 0 {
		fmt.Println("\n[Peringatan] Belum ada data tempat yang terdaftar.")
		return
	}

	fmt.Println("\n--- DAFTAR CO-WORKING SPACE ---")
	for i := 0; i < len(daftarSpace); i++ {
		space := daftarSpace[i]
		statusWiFi := "Tidak Ada"
		if space.HasWiFi {
			statusWiFi = "Tersedia"
		}
		fmt.Printf("ID: %d | %s (%s) | Harga: Rp%.0f/jam | Rating: %.1f | WiFi: %s\n",
			space.ID, space.Nama, space.Lokasi, space.HargaSewa, space.Rating, statusWiFi)
	}
}

func tambahSpaceBaru() {
	var namaInput, lokasiInput string
	var hargaInput, ratingInput float64
	var wifiInput string
	var wifiBool bool

	fmt.Println("\n--- FORM TAMBAH DATA CO-WORKING SPACE ---")
	fmt.Print("Nama Tempat         : ")
	fmt.Scanln(&namaInput)
	fmt.Print("Lokasi/Kota         : ")
	fmt.Scanln(&lokasiInput)
	fmt.Print("Harga Sewa (per jam): ")
	fmt.Scanln(&hargaInput)
	fmt.Print("Berikan Rating (1-5): ")
	fmt.Scanln(&ratingInput)
	fmt.Print("Apakah ada WiFi? (y/n): ")
	fmt.Scanln(&wifiInput)

	if wifiInput == "y" || wifiInput == "Y" {
		wifiBool = true
	}

	totalData++
	dataBaru := model.CoWorkingSpace{
		ID:        totalData,
		Nama:      namaInput,
		Lokasi:    lokasiInput,
		HargaSewa: hargaInput,
		Rating:    ratingInput,
		HasWiFi:   wifiBool,
	}

	daftarSpace = append(daftarSpace, dataBaru)
	fmt.Println("\n[Sukses] Data co-working space baru berhasil disimpan!")
}

func menuSorting() {
	if len(daftarSpace) == 0 {
		fmt.Println("\n[Peringatan] Tidak ada data yang bisa diurutkan.")
		return
	}

	var subPilihan int
	fmt.Println("\n--- PILIHAN ALGORITMA PENGURUTAN ---")
	fmt.Println("1. Urutkan Harga Sewa Termurah (Selection Sort Ascending)")
	fmt.Println("2. Urutkan Rating Tertinggi (Insertion Sort Descending)")
	fmt.Print("Pilih opsi (1-2): ")
	fmt.Scanln(&subPilihan)

	if subPilihan == 1 {
		alpro.SelectionSortHarga(daftarSpace)
		fmt.Println("\n[Sukses] Data berhasil diurutkan berdasarkan harga termurah!")
		tampilkanSemuaSpace()
	} else if subPilihan == 2 {
		alpro.InsertionSortRating(daftarSpace)
		fmt.Println("\n[Sukses] Data berhasil diurutkan berdasarkan rating tertinggi!")
		tampilkanSemuaSpace()
	} else {
		fmt.Println("\nPilihan opsi tidak valid.")
	}
}

// PROSEDUR MODULAR: Sub-menu khusus menangani searching
func menuSearching() {
	if len(daftarSpace) == 0 {
		fmt.Println("\n[Peringatan] Tidak ada data yang bisa dicari.")
		return
	}

	var subPilihan int
	fmt.Println("\n--- PILIHAN ALGORITMA PENCARIAN ---")
	fmt.Println("1. Cari Berdasarkan Lokasi/Kota (Sequential Search)")
	fmt.Println("2. Cari Berdasarkan Nama Tempat Persisi (Binary Search)")
	fmt.Print("Pilih opsi (1-2): ")
	fmt.Scanln(&subPilihan)

	if subPilihan == 1 {
		var lokasiCari string
		fmt.Print("Masukkan nama lokasi/kota yang dicari: ")
		fmt.Scanln(&lokasiCari)

		hasil := alpro.SequentialSearchLokasi(daftarSpace, lokasiCari)
		if len(hasil) == 0 {
			fmt.Println("\n[Hasil] Tidak ada tempat di lokasi tersebut.")
		} else {
			fmt.Println("\n--- HASIL PENCARIAN SEKUENSIAL ---")
			for _, space := range hasil {
				fmt.Printf("%s di %s | Harga: Rp%.0f/jam\n", space.Nama, space.Lokasi, space.HargaSewa)
			}
		}
	} else if subPilihan == 2 {
		var namaCari string
		fmt.Print("Masukkan Nama Tempat secara tepat: ")
		fmt.Scanln(&namaCari)

		idx := alpro.BinarySearchNama(daftarSpace, namaCari)
		if idx == -1 {
			fmt.Println("\n[Hasil] Tempat dengan nama tersebut tidak ditemukan.")
		} else {
			fmt.Println("\n--- HASIL PENCARIAN BINARY (KETEMU) ---")
			space := daftarSpace[idx]
			fmt.Printf("ID: %d | %s | Lokasi: %s | Harga: Rp%.0f/jam | Rating: %.1f\n",
				space.ID, space.Nama, space.Lokasi, space.HargaSewa, space.Rating)
		}
	} else {
		fmt.Println("\nPilihan opsi tidak valid.")
	}
}

func ubahSpace() {
	var idCari int
	fmt.Println("\n--- FORM UBAH DATA CO-WORKING SPACE ---")
	fmt.Print("Masukkan ID Co-Working Space yang ingin diubah: ")
	fmt.Scanln(&idCari)

	indexDitemukan := -1
	for i := 0; i < len(daftarSpace); i++ {
		if daftarSpace[i].ID == idCari {
			indexDitemukan = i
			break
		}
	}

	if indexDitemukan == -1 {
		fmt.Println("[Error] Data dengan ID tersebut tidak ditemukan!")
		return
	}

	var namaInput, lokasiInput string
	var hargaInput, ratingInput float64
	var wifiInput string

	fmt.Printf("\nData Lama: %s (%s)\n", daftarSpace[indexDitemukan].Nama, daftarSpace[indexDitemukan].Lokasi)
	fmt.Print("Masukkan Nama Baru         : ")
	fmt.Scanln(&namaInput)
	fmt.Print("Masukkan Lokasi Baru       : ")
	fmt.Scanln(&lokasiInput)
	fmt.Print("Masukkan Harga Sewa Baru   : ")
	fmt.Scanln(&hargaInput)
	fmt.Print("Masukkan Rating Baru (1-5) : ")
	fmt.Scanln(&ratingInput)
	fmt.Print("Apakah ada WiFi? (y/n)     : ")
	fmt.Scanln(&wifiInput)

	daftarSpace[indexDitemukan].Nama = namaInput
	daftarSpace[indexDitemukan].Lokasi = lokasiInput
	daftarSpace[indexDitemukan].HargaSewa = hargaInput
	daftarSpace[indexDitemukan].Rating = ratingInput
	if wifiInput == "y" || wifiInput == "Y" {
		daftarSpace[indexDitemukan].HasWiFi = true
	} else {
		daftarSpace[indexDitemukan].HasWiFi = false
	}

	fmt.Println("\n[Sukses] Data berhasil diperbarui!")
}

func hapusSpace() {
	var idCari int
	fmt.Println("\n--- FORM HAPUS DATA CO-WORKING SPACE ---")
	fmt.Print("Masukkan ID Co-Working Space yang ingin dihapus: ")
	fmt.Scanln(&idCari)

	indexDitemukan := -1
	for i := 0; i < len(daftarSpace); i++ {
		if daftarSpace[i].ID == idCari {
			indexDitemukan = i
			break
		}
	}

	if indexDitemukan == -1 {
		fmt.Println("[Error] Data dengan ID tersebut tidak ditemukan!")
		return
	}

	daftarSpace = append(daftarSpace[:indexDitemukan], daftarSpace[indexDitemukan+1:]...)
	fmt.Println("\n[Sukses] Data co-working space berhasil dihapus!")
}
