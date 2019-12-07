package main

import (
	"io"
	"strings"
	"testing"
)

func TestDay3PartAEg1(t *testing.T) {
	reader := strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	doTest(reader, false, 159, t)
}

func TestDay3PartAEg2(t *testing.T) {
	reader := strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	doTest(reader, false, 135, t)
}

func TestDay3PartBEg1(t *testing.T) {
	reader := strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	doTest(reader, true, 610, t)
}

func TestDay3PartBEg2(t *testing.T) {
	reader := strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	doTest(reader, true, 410, t)
}

func doTest(reader io.Reader, partB bool, expected int, t *testing.T) {
	actual := run(reader, partB)

	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
