package main

import (
	"testing"
)

var ages []PersonAges

// init data before test dimulai
func seedAges() {
	ages = []PersonAges{
		{"person:12", 24},
		{"person:134", 26},
		{"person:112", 23},
		{"person:1412", 23},
		{"person:421", 23},
		{"person:415", 26},
	}
}

// input nya array of PersonAges
// expected value: map[21:2 22:1 28:2]
func Test_GetAgeDistr(t *testing.T) {
	seedAges()
	gots := GetAgeDistr(ages)
	expected := map[int]occurence{
		24: 1,
		23: 3,
		26: 2,
	}
	if len(gots) != len(expected) {
		t.Errorf("fail because len of data: got %v, expected %v", len(gots), len(expected))
	}
	for key, _ := range gots {
		if gots[key] != expected[key] {
			t.Errorf("got %v, expected %v", gots[key], expected[key])
		}
	}

}

// input : tampilkan people dengan umur 23 only
// expected value: [{person:12 21} {person:32 21}]
func Test_FindByAge(t *testing.T) {
	seedAges()
	findAge := 23
	gots := FindByAge(ages, findAge)

	expected := map[string]int{
		"person:112":  23,
		"person:1412": 23,
		"person:421":  23,
	}

	if len(gots) != len(expected) {
		t.Errorf("fail because len of data: got %v, expected %v", len(gots), len(expected))
	}

	// jika masih kosong, ini tidak akan pernah di run
	for _, val := range gots {
		if _, ok := expected[val.id]; !ok {
			t.Errorf("%v is not found on expected value", val.id)
		}
	}

}

// input : tampilkan people dengan umur 50 only
// expected value: error not found
// func Test_FindByAge_notfound(t *testing.T) {
// 	// findAge := 23
// 	// got := FindByAge(ages, findAge)
// 	// // collection
// 	// expected := ""

// 	// t.Errorf("got %v, expected %v", got, expected)
// }
