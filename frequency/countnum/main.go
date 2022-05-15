package main

import "fmt"

type frequencies map[int32]int32

// arr to hashmap
func count(arr []int32) frequencies {
	freq := make(frequencies)
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
