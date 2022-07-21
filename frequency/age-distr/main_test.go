package main

import "testing"

func Test_GetAgeDistr(t *testing.T) {
	ages := []PersonAges{
		{"person:12", 24},
		{"person:134", 26},
		{"person:112", 23},
		{"person:1412", 23},
		{"person:421", 23},
		{"person:415", 26},
	}
	got := GetAgeDistr(ages)
	expected := map[int]int{
		24: 1,
		23: 3,
		26: 2,
	}
	for key, _ := range got {
		if got[key] != expected[key] {
			t.Errorf("got %v, expected %v", got[key], expected[key])
		}
	}

}
