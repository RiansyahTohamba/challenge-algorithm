package main

import (
	"fmt"
	"testing"
)

func Test_diff_val_for_rank_1and2(t *testing.T) {
	list := []int{5, 9, 7, 11, 10, 20, 30}
	fmt.Println()
	got := findMaxSum(list)
	expected := 20 + 30
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func Test_same_rank1_2(t *testing.T) {
	list := []int{5, 9, 7, 11, 10, 20, 20}
	fmt.Println()
	got := findMaxSum(list)

	// jadi harus distinct then baru dijumlahkan

	expected := 20 + 11
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
