package main

import "fmt"

/*
given, collection of lilin.
dapatkan lilin tertinggi.
hitung total lilin tertinggi pada kumpulan lilin.

how to solve it?
	bagaiman mencari nilai maksimum pada array?
	caranya adalah membandingkan satu persatu
{1,2,1,5,3}
bubble comparison
1 vs 2
1 vs 1
1 vs 5
1 vs 3

2 vs 1
2 vs 5
etc
...
bagaimana hashmap solve bubble?
adalah dengan menetapkan nilai lilin terbesar terlebih dahulu

awalnya 0, lalu dia akan dibandingkan dengan candles lainnya.
jika tallestcandle lebih kecil dari candle yang akan diambil maka dia tidak lagi menjadi tallest candle
*/
func main() {
	// padahal ada 2 lilin 5 tetapi hanya dihitung 1
	// bagaimana jika ada lilin yang sama?
	arr := []int32{2, 3, 4, 5, 5, 1}
	fmt.Println(birthdayCakeCandles(arr))
	arr = []int32{2, 3, 4, 6, 6, 6, 4, 4}
	fmt.Println(birthdayCakeCandles(arr))
}

func birthdayCakeCandles(candles []int32) int32 {

	var tallestCandle int32
	numTallestCandle := make(map[int32]int32)

	for _, candle := range candles {
		// swapping ranking hanya terjadi sekali, ga mungkin berkali=kali
		if tallestCandle < candle || tallestCandle == candle {
			// swap
			tallestCandle = candle
			// counting + 1
			numTallestCandle[candle]++
		}

	}
	return numTallestCandle[tallestCandle]
}
