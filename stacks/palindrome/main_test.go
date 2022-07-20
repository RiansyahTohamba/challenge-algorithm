package main

import "testing"

func TestPalindromeTrue(t *testing.T) {
	word := "anna"
	expected := true
	got := isPalindrome(word)
	if got != expected {
		t.Errorf("want '%v', got '%v' ", expected, got)
	}
}

func TestPalindromeFalse(t *testing.T) {
	word := "bionda"
	expected := false
	got := isPalindrome(word)
	if got != expected {
		t.Errorf("want '%v', got '%v' ", expected, got)
	}
}
