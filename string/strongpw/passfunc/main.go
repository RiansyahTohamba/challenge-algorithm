package main

import (
	"fmt"
	"unicode"
)

func isContain(password string, check func(rune) bool) bool {
	for _, pw := range password {
		if check(pw) {
			return true
		}
	}

	return false
}

func main() {
	password := "aian"
	check := unicode.IsUpper
	result := isContain(password, check)
	fmt.Println(result)
}
