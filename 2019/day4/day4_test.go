package main

import (
	"testing"
)

func TestPartBExamples(t *testing.T) {
	doTest(111111, true, false, t)
	doTest(333445, true, true, t)
	doTest(334444, true, true, t)
	doTest(112233, true, true, t)
	doTest(123444, true, false, t)
	doTest(111122, true, true, t)
}

func TestPartAExamples(t *testing.T) {
	doTest(111111, false, true, t)
	doTest(223450, false, false, t)
	doTest(123789, false, false, t)
}

func doTest(v int, partB bool, expected bool, t *testing.T) {
	res := evaluateCandidate(v, partB)
	if res != expected {
		t.Errorf("evaluate %d = %t; want %t", v, res, expected)
	}
}
