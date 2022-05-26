package main

import "fmt"

func arrToMap(items []int32) map[int32]int32 {
	freq := make(map[int32]int32)
	for _, val := range items {
		freq[val]++
	}
	return freq
}

func FindItem(items []int32) int32 {
	// convert array to map
	freq := arrToMap(items)

	var max int32 = freq[items[0]]
	var key int32 = items[0]

	for idx, val := range freq {
		if max <= val && key > idx {
			max = val
			key = idx
		}
	}
	return key
}

func main() {
	// pada case ini, 1 muncul pertama,
	// maka max_frequencies=2
	// maka max=2 tidak akan tergantikan karena 2 < 2 is false
	// maka key yang dipilih selalu key yang pertama kali muncul

	items := []int32{2, 2, 1, 1}

	mostId := FindItem(items)
	fmt.Println(mostId)
	// isTrue := (2 <= 2) && (2 > 1)
	// fmt.Println(isTrue)
}
