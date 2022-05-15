package main

import "fmt"

// https://www.hackerrank.com/challenges/migratory-birds/problem?isFullScreen=true

// int: the lowest type id of the most frequently sighted birds

func migratoryBirds(arr []int32) int32 {
	var max int32
	// Write your code here
	freq := make(map[int32]int32)
	for _, el := range arr {
		if el > max {
			max = el
			freq[max] = 1
		} else if max == el {
			freq[max]++
		}
	}
	// return id terkecil

	// masalahnya bagaimana jika ada 2 nilai yang sama besar?
	// lalu bagaimana cara mengambil nilai terkecil?
	return freq[max]
}

func main() {
	birds := []int32{3, 3, 3, 3, 1, 1, 1, 4, 5}
	fmt.Println(migratoryBirds(birds))
}
