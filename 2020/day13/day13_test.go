package main

import "testing"

func TestPartBEg(t *testing.T) {
	ans := partB([]bus{{7, 0}, {13, 1}, {59, 4}, {31, 6}, {19, 7}})
	checkEq(ans, 1068781, t)
}

func TestFindInverse1(t *testing.T) {

	ans := findInverse(5, 3)
	checkEq(ans, 2, t)
}

func TestFindInverse2(t *testing.T) {

	ans := findInverse(3, 5)
	checkEq(ans, 2, t)
}

func checkEq(actual, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}
}
