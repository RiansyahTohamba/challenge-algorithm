package main

import "fmt"

// relation between array and hashmap
func main() {
	names := []string{"rian", "bio", "tohamba", "fathan"}
	courses := map[string]string{
		"rian":   "math",
		"bio":    "physics",
		"fathan": "physics",
	}
	join(names, courses)
}

// top := len(stack) - 1
// bracketFromStack := stack[top]
// fmt.Println("dari stack adanya ini")
// fmt.Println(string(bracketFromStack))

// fmt.Println("value dari DB ini")
// fmt.Println(string(bracketDB[bracketFromStack]))

// fmt.Println("key yang datang dari string")
// fmt.Println(string(key))

func join(names []string, courses map[string]string) {
	top := len(names) - 1
	fromStack := names[top]
	fmt.Println(fromStack)

	res := courses[fromStack]
	fmt.Println(res)
}
