package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	// alokasi array dengan len=6 (0-5)
	mapBird := make([]int, 6)
	// index pada array mapBird adalah birdId
	// value nya adalah total kemunculan birdId pada set
	for _, birdId := range arr {
		mapBird[birdId]++
	}
	// fmt.Println(mapBird)
	// inisiasi maxFrequencies dengan nilai pertama dari arr
	maxFreq := mapBird[arr[0]]

	for _, freq := range mapBird {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	// temukan minIdx
	for minIdx, freq := range mapBird {
		if freq == maxFreq {
			return int32(minIdx)
		}
	}
	return 1
}

func main() {
	// case dimana id burung harus masuk dalam range (1..5) = {1,2,3,4,5}
	// tidak boleh diluar itu (diatas 5,6,...)

	birds := []int32{1, 1, 1, 2, 2, 2, 3, 4}

	minId := migratoryBirds(birds)

	fmt.Printf("\nid terkecil dan tersering muncul= %d \n", minId)
}
