package main

import (
	"fmt"
	"math"
	"unicode"
)

func isContainChar(password string, isContain func(rune) bool) bool {
	for _, pw := range password {
		if isContain(pw) {
			return false
		}
	}
	return true
}

func isSpecialChar(charpw rune) bool {
	var special_characters string = "!@#$%^&*()-+"
	for _, sc := range special_characters {
		if charpw == sc {
			return true
		}
	}
	return false
}
func minimumNumber(n int32, password string) int32 {
	var minLength int32 = 6

	countWord := 0
	// 2. It contains at least one digit.
	if isContainChar(password, unicode.IsNumber) {
		countWord++
	}
	// 3. It contains at least one lowercase English character.
	if isContainChar(password, unicode.IsLower) {
		countWord++
	}
	// 4. It contains at least one uppercase English character.
	if isContainChar(password, unicode.IsUpper) {
		countWord++
	}
	if isContainChar(password, isSpecialChar) {
		countWord++
	}

	// 2 - 5 : check contain one digit kah?

	// 1. Its length is at least 6.
	lengthDiff := float64(minLength - n)

	count64 := float64(countWord)

	maxVal := int32(math.Max(count64, lengthDiff))
	return maxVal
}

func main() {
	noUppSpec := "abcdexa1!1A"
	fmt.Println(minimumNumber(8, noUppSpec))
}
