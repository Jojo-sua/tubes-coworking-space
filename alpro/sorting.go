package alpro

import (
	"strings"
	"tubes_coworking/model"
)

// 1. SELECTION SORT (Ascending: Murah -> Mahal)
func SelectionSortHarga(arr []model.CoWorkingSpace) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		posisiMin := i

		// Cari data yang harganya paling murah di sisa array
		for j := i + 1; j < n; j++ {
			if arr[j].HargaSewa < arr[posisiMin].HargaSewa {
				posisiMin = j
			}
		}

		// Tukar posisi data pakai variabel penampung (temp)
		temp := arr[i]
		arr[i] = arr[posisiMin]
		arr[posisiMin] = temp
	}
}

// 2. INSERTION SORT (Descending: Rating Tinggi -> Rendah)
func InsertionSortRating(arr []model.CoWorkingSpace) {
	n := len(arr)

	for i := 1; i < n; i++ {
		target := arr[i]
		j := i - 1

		// Geser data ke kanan selama ratingnya lebih kecil dari target
		for j >= 0 && arr[j].Rating < target.Rating {
			arr[j+1] = arr[j]
			j--
		}

		// Selipkan data target di posisi yang pas
		arr[j+1] = target
	}
}

// 3. SEQUENTIAL SEARCH (Cari lokasi yang mirip/mengandung teks)
func SequentialSearchLokasi(arr []model.CoWorkingSpace, kataKunci string) []model.CoWorkingSpace {
	var hasil []model.CoWorkingSpace
	cari := strings.ToLower(kataKunci)

	for i := 0; i < len(arr); i++ {
		lokasiSistem := strings.ToLower(arr[i].Lokasi)

		// Kalau teks lokasi mengandung kata kunci, masukkan ke hasil
		if strings.Contains(lokasiSistem, cari) {
			hasil = append(hasil, arr[i])
		}
	}
	return hasil
}

// 4. BINARY SEARCH (Cari nama tempat yang presisi)
func BinarySearchNama(arr []model.CoWorkingSpace, namaCari string) int {
	n := len(arr)

	// Urutkan data berdasarkan nama dulu (Syarat mutlak Binary Search)
	for i := 0; i < n-1; i++ {
		posisiMin := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(arr[j].Nama) < strings.ToLower(arr[posisiMin].Nama) {
				posisiMin = j
			}
		}
		temp := arr[i]
		arr[i] = arr[posisiMin]
		arr[posisiMin] = temp
	}

	// Mulai pencarian bagi dua (Binary Search)
	kiri := 0
	kanan := n - 1
	cari := strings.ToLower(namaCari)

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaTengah := strings.ToLower(arr[tengah].Nama)

		if namaTengah == cari {
			return tengah // Ketemu, kembalikan indeksnya
		} else if namaTengah < cari {
			kiri = tengah + 1 // Cari di sebelah kanan
		} else {
			kanan = tengah - 1 // Cari di sebelah kiri
		}
	}
	return -1 // Tidak ketemu
}
