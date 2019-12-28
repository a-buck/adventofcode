package main

import "testing"

func TestEg1Phase1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ans := ftt(input, 1)
	checkSliceEq(ans, []int{4, 8, 2, 2, 6, 1, 5, 8}, t)
}

func TestEg1Phase2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ans := ftt(input, 2)
	checkSliceEq(ans, []int{3, 4, 0, 4, 0, 4, 3, 8}, t)
}

func TestPatternIndex(t *testing.T) {
	// elem 0
	patternTest(0, 0, 1, t)
	patternTest(1, 0, 0, t)
	patternTest(2, 0, -1, t)
	patternTest(3, 0, 0, t)
	patternTest(4, 0, 1, t)
	patternTest(5, 0, 0, t)
	patternTest(6, 0, -1, t)
	patternTest(7, 0, 0, t)

	// elem 1
	patternTest(0, 1, 0, t)
	patternTest(1, 1, 1, t)
	patternTest(2, 1, 1, t)
	patternTest(3, 1, 0, t)
	patternTest(4, 1, 0, t)
	patternTest(5, 1, -1, t)
	patternTest(6, 1, -1, t)
	patternTest(7, 1, 0, t)

	// elem 2
	patternTest(0, 2, 0, t)
	patternTest(1, 2, 0, t)
	patternTest(2, 2, 1, t)
	patternTest(3, 2, 1, t)
	patternTest(4, 2, 1, t)
	patternTest(5, 2, 0, t)
	patternTest(6, 2, 0, t)
	patternTest(7, 2, 0, t)
}

func patternTest(i, e int, expected int, t *testing.T) {
	actual := pattern(i, e)
	if actual != expected {
		t.Errorf("pattern(i=%d, e=%d)=%d, expected %d", i, e, actual, expected)
	}
}

func checkSliceEq(actual []int, expected []int, t *testing.T) {
	if len(actual) != len(expected) {
		t.Errorf("got: %+v, expected: %+v", actual, expected)
	}

	for i, a := range actual {
		if a != expected[i] {
			t.Errorf("got: %+v, expected: %+v", actual, expected)
		}
	}
}
