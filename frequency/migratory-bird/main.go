package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	// alokasi array dengan len=6 (0-5)
	counts := make([]int, 6)
	// index pada array counts adalah birdId
	// value nya adalah total kemunculan birdId pada set
	for _, birdId := range arr {
		counts[birdId] += 1
	}
	fmt.Println(counts)
	// mencari maksimum frequency
	maxFreq := counts[0]
	for _, freq := range counts {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	// temukan minIdx
	for minIdx, freq := range counts {
		if freq == maxFreq {
			return int32(minIdx)
		}
	}
	return 1
}

func main() {
	// case dimana
	// id burung harus masuk dalam set = {1,2,3,4,5}
	// tidak boleh diluar itu

	birds := []int32{1, 1, 1, 2, 2, 2, 3, 4}

	minId := migratoryBirds(birds)

	fmt.Printf("\nid terkecil dan tersering muncul= %d \n", minId)
}
