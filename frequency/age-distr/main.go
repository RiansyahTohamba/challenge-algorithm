package main

import "fmt"

type PersonAges struct {
	// member
	id string
	// score
	age int
}

func main() {
	ages := []PersonAges{{id: "person:12", age: 21},
		{id: "person:88", age: 28}, {id: "person:32", age: 21},
		{id: "person:77", age: 22}, {id: "person:21", age: 28}}

	fmt.Println(GetAgeDistr(ages))
	fmt.Println(FilterByAge(ages, 21))
}

func GetAgeDistr(ages []PersonAges) map[int]int {
	ageDistr := map[int]int{}
	for _, ag := range ages {
		ageDistr[ag.age]++
	}
	return ageDistr
}

func FilterByAge(ages []PersonAges, age int) []PersonAges {
	peopleByAge := make(map[int][]PersonAges)
	for _, person := range ages {
		peopleByAge[person.age] = append(peopleByAge[person.age], person)
	}
	return peopleByAge[age]
}
