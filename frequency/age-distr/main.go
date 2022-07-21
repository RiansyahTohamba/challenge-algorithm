package main

import "fmt"

type occurence int

type PersonAges struct {
	// member
	id string
	// score
	age int
}

func main() {
	expected := map[string]int{
		"person:112":  23,
		"person:1412": 23,
		"persosd":     23,
	}
	for k, v := range expected {
		fmt.Println(k)
		fmt.Println(v)
	}
}

// map[int]occurence = map[age]int
// expected value: map[21:2 22:1 28:2]
func GroupByAge(ages []PersonAges) map[int]occurence {
	agesDistr := make(map[int]occurence)
	for _, p := range ages {
		agesDistr[p.age]++
	}
	return agesDistr
}

func FindByAge(ages []PersonAges, age int) []PersonAges {
	prsAges := make([]PersonAges, 0)
	for _, p := range ages {
		if p.age == age {
			prsAges = append(prsAges, p)
		}
	}
	return prsAges
}
