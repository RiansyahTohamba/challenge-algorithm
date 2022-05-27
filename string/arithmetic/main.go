package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// ?
func count(match []string) int {
	operator := match[2]

	nums_1, _ := strconv.Atoi(match[1])
	nums_2, _ := strconv.Atoi(match[3])

	switch {
	case operator == "+":
		return nums_1 + nums_2
	case operator == "-":
		return nums_1 - nums_2
	case operator == "*":
		return nums_1 * nums_2
	case operator == "/":
		return nums_1 / nums_2
	}
	return 0
}

func main() {
	rumus := "2 + 2"

	reg, _ := regexp.Compile(`(\d+)(?:\s*)([\+\-\*\/])(?:\s*)(\d+)`)
	// (?:\s*) exclude space?
	// ketika di compile, yang diambil hanya (\d+),([\+\-\*\/]),(\d+)
	match := reg.FindStringSubmatch(rumus)

	res := count(match)
	fmt.Println(res)
}
