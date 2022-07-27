package main

import "fmt"

type Empl struct {
	f string
	l string
}
type Kar struct {
	Empl
}

func (e Empl) fullname() string {
	return e.f + " " + e.l
}
func main() {
	m := Kar{Empl: Empl{"na", "la"}}
	fmt.Println(m.Empl.fullname())

	fmt.Println(m.f + " " + m.l)
	fmt.Println(m.fullname())
}
