```go
package main

import "fmt"

// https://www.hackerrank.com/challenges/migratory-birds/problem?isFullScreen=true

// int: the lowest type id of the most frequently sighted birds

func migratoryBirds(arr []int32) int32 {
	var occur int32
	distribution := make(map[int32]int32)
	for _, el := range arr {
		if el < occur {
			occur = el
			distribution[occur] = 1
		} else if occur == el {
			distribution[occur]++
		}
	}
	// return id terkecil

	// masalahnya bagaimana jika ada 2 nilai yang sama besar?
	// lalu bagaimana cara mengambil nilai terkecil?
	return distribution[occur]
}

// case dimana
// bilangan terkecil hanya punya 2 kali muncul
// bilangan paling sering muncul tapi muncul 3 kali
// jadi ini bukan tentang mencari nilai terkecil
// tetapi mencari frekuensi terbesar terlebih dahulu
func main() {
	// case dimana
	birds := []int32{3, 3, 3, 1, 1, 4, 5}
	fmt.Println(migratoryBirds(birds))
}
