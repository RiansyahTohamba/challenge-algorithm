package main

import (
	"fmt"
	"unicode"
)

// hitung total word pada s
func camelCase(s string) int32 {
	// bagaimana mengecek upercase vs lowercase?
	// isLowerCase
	var countWord int32 = 1

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			countWord++
		}
	}
	return countWord
}

func main() {
	fmt.Println(camelCase("oneTwoThree"))

	fmt.Println(unicode.IsUpper('p'))

	fmt.Println(unicode.IsUpper('P'))
}
