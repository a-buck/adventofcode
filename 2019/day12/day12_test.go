package main

import "testing"

func TestLcm(t *testing.T) {
	res := lcm(4, 6)
	checkEq(res, 12, t)
}

func checkEq(actual, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
