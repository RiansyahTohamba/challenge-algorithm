package main

import (
	"fmt"
	"regexp"
)

// ("    United Kingdom    ")
// ("united-kingdom")
// ""
var jawaban = "United Kingdom"

func main() {
	userInput := "    United    Kingdom    "
	reg, _ := regexp.Compile(`(\w+)(?:\s*)(\w+)`)
	// (?:\s*) exclude space?
	// ketika di compile, yang diambil hanya (\d+),([\+\-\*\/]),(\d+)
	match := reg.FindStringSubmatch(userInput)
	for _, v := range match {
		fmt.Println(v)
	}
	fmt.Println(match[0])

}
