package alpro

import (
	"strings"
	"tubes_coworking/model"
)

// Sequential Search untuk mencari Lokasi (Partial Match / Kemiripan Teks)
func SequentialSearchLokasi(arr []model.CoWorkingSpace, kataKunci string) []model.CoWorkingSpace {
	var hasil []model.CoWorkingSpace

	kunciCari := strings.ToLower(kataKunci)

	for i := 0; i < len(arr); i++ {
		lokasiTempat := strings.ToLower(arr[i].Lokasi)
		// Cek apakah kata kunci ada di dalam teks lokasi tempat
		if strings.Contains(lokasiTempat, kunciCari) {
			hasil = append(hasil, arr[i])
		}
	}
	return hasil
}

// Binary Search untuk mencari Nama Tempat
func BinarySearchNama(arr []model.CoWorkingSpace, namaCari string) int {
	// Sorting
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(arr[j].Nama) < strings.ToLower(arr[minIdx].Nama) {
				minIdx = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = temp
	}

	// Proses Binary Search
	low := 0
	high := n - 1
	kunciCari := strings.ToLower(namaCari)

	for low <= high {
		mid := (low + high) / 2
		namaMid := strings.ToLower(arr[mid].Nama)

		if namaMid == kunciCari {
			return mid // Mengembalikan index jika ketemu
		} else if namaMid < kunciCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Mengembalikan -1 jika tidak ketemu
}
