package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Lisa", age: 21},
		{name: "Sam", age: 28}, {name: "Doni", age: 21},
		{name: "Joe", age: 22}, {name: "Rani", age: 28}}
	fmt.Println(ageDistr(people))
	fmt.Println(FilterByAge(people, 21))
}

func ageDistr(people []Person) map[int]int {
	ageDistr := map[int]int{}
	for _, person := range people {
		ageDistr[person.age]++
	}
	return ageDistr
}

func FilterByAge(people []Person, age int) []Person {
	peopleByAge := make(map[int][]Person)
	for _, person := range people {
		peopleByAge[person.age] = append(peopleByAge[person.age], person)
	}
	return peopleByAge[age]
}
