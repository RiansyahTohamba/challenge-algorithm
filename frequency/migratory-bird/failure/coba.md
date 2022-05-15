package main

import "fmt"

func NotmigratoryBirds(arr []int32) int32 {
	int max = 0;
	int id = 0;
	for (int i = 0; i < 5; i++) {
		if (birds[i] > max) {
			max = birds[i];
			id = i+1;
		}
	}
	return id


}

func NotmigratoryBirds(arr []int32) int32 {
	counts := make(map[int32]int32)
	var minKey int32 = 1
	var max int32 = 1

	for _, el := range arr {
		counts[el]++
	}
	
	// 3 dan 1 pada case dibawah
	// tapi yang muncul hanya 3?
	// karena 3 > 3, false
	//

	// mencari frekuensi tertinggi
	for key, val := range counts {
		// algoritma ini tidak memenuhi case dimana 1 adalah id terkecil tetapi frekuensi terkecil juga
		if val > max {
			max = val
			minKey = key
			fmt.Println("id ini sering muncul")
			fmt.Println(key)
			// max = counts[key]
			// minKey = key
		}
	}

	// 1. bagaimana cara mengetahui bilangan terkecil pada set?
	// 2. bagaimana cara mengetahui frekuensi tertinggi?

	// kalau koding, cari jawaban problem nomor 1 dulu atau yang 2 dulu?
	for key, val := range counts {
		fmt.Printf("%d = %d kali \n", key, val)
	}

	return minKey
}

func main() {
	// case dimana
	// id burung harus masuk dalam set = {1,2,3,4,5}
	// tidak boleh diluar itu
	birds := []int32{3, 3, 3, 2, 2, 2, 1, 4, 5}

	minId := migratoryBirds(birds)

	fmt.Printf("\nid terkecil dan tersering muncul= %d \n", minId)
}
