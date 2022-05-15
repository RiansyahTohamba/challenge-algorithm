package main

// mencari frekuensi untuk nilai maksimum yang ada pada set

// tidak perlu mencatat frekuensi semua bilangan
// cukup berfokus kepada nilai maksimum pada array
func getMax(arr []int) int {
	// terdapat 2 variable
	// variable untuk menampung nilai maksimum (sementara)
	var max int
	// variabel untuk menyimpan key-value antara max_val-freq
	freq := make(map[int]int)

	for _, el := range arr {
		if el > max {
			max = el
			freq[max] = 1
		} else if max == el {
			// jika ketemu lagi dengan elemen yang sama dengan nilai maksimum maka tambah nilainya
			freq[max]++
		}
	}
	return freq[max]
}
