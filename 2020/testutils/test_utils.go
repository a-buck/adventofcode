package testutils

import "testing"

// CheckStringsEq check 2 strings are equal
func CheckStringsEq(actual, expected string, t *testing.T) {
	if string(actual) != string(expected) {
		t.Errorf("Got %s, expected %s", actual, expected)
	}
}

// CheckIntsEq check 2 ints are equal
func CheckIntsEq(actual, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}
}
