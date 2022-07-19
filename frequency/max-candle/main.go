package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 2, 4, 2, 3}
	max := getFreqOfMaxEl(arr)
	fmt.Println(max)
}

// mencari frekuensi untuk element dengan nilai tertinggi pada set

// tidak perlu mencatat frekuensi semua bilangan
// cukup berfokus kepada nilai maksimum pada array
func getFreqOfMaxEl(arr []int) int {
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
