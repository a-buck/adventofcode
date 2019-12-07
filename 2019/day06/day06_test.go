package main

import (
	"strings"
	"testing"
)

func TestPartAEg(t *testing.T) {
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	runTest(input, false, 42, t)
}

func TestPartBEg(t *testing.T) {
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	runTest(input, true, 4, t)
}

func runTest(input string, partB bool, expected int, t *testing.T) {
	actual, err := run(strings.NewReader(input), partB)

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
