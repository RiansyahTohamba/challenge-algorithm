package main

import "fmt"

// arr to hashmap
func count(arr []int32) map[int32]int32 {
	freq := make(map[int32]int32)
	for _, el := range arr {
		freq[el]++
	}
	return freq
}

func main() {
	arr := []int32{1, 2, 1, 1, 3, 3, 3, 4}
	fmt.Println("map[num:freq]")
	for key, val := range count(arr) {
		fmt.Printf("%d= %d kali \n", key, val)
	}
}
