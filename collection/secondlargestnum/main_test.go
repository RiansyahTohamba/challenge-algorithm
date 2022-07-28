package main

import (
	"testing"
)

func Test_diff_val_for_rank_1and2(t *testing.T) {
	list := []int{5, 9, 7, 11, 10, 20, 30}
	r1, r2 := getFirstSecondRank(list)
	got := r1 + r2
	ranking_1 := 30
	ranking_2 := 20
	expected := ranking_1 + ranking_2
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func Test_same_rank1_2(t *testing.T) {
	list := []int{5, 9, 7, 11, 10, 20, 20}

	r1, r2 := getFirstSecondRank(list)
	got := r1 + r2

	ranking_1 := 20
	ranking_2 := 11
	expected := ranking_1 + ranking_2

	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
