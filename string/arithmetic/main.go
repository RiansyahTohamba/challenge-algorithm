package main

import (
	"fmt"
	"unicode"
)

// ?
func count(nums []int, operator rune) int {
	switch {
	case operator == '+':
		return nums[0] + nums[1]
	case operator == '-':
		return nums[0] - nums[1]
	case operator == '*':
		return nums[0] * nums[1]
	case operator == '/':
		return nums[0] / nums[1]
	}
	return 0
}

func isOperator(char rune) bool {
	var ops string = "+-*/"
	for _, op := range ops {
		if char == op {
			return true
		}
	}

	return false
}

func main() {
	rumus := "2 + 2"
	nums := make([]int, 2)
	var arithmOp rune
	// find digit, set it as first element
	// find operator, set it as second element
	// find second digit, set it as third element
	// /(\d+)(?:\s*)([\+\-\*\/])(?:\s*)(\d+)/
	for _, ch := range rumus {
		if unicode.IsDigit(ch) {
			// nah masalahnya jika 9000 +90, wkwk, bakalan false nih
			nums = append(nums, int(ch))
		}
		if isOperator(ch) {
			arithmOp = ch
		}
	}

	res := count(nums, arithmOp)
	fmt.Println(res)
}
