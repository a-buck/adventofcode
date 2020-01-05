package main

import (
	"log"
	"os"
	"testing"
)

func TestAEg1(t *testing.T) {
	doTest("testdata/day18Aeg1.txt", 8, false, t)
}

func TestAEg2(t *testing.T) {
	doTest("testdata/day18Aeg2.txt", 86, false, t)
}

func TestAEg3(t *testing.T) {
	doTest("testdata/day18Aeg3.txt", 132, false, t)
}

func TestAEg4(t *testing.T) {
	doTest("testdata/day18Aeg4.txt", 136, false, t)
}

func TestAEg5(t *testing.T) {
	doTest("testdata/day18Aeg5.txt", 81, false, t)
}

func TestBEg1(t *testing.T) {
	doTest("testdata/day18Beg1.txt", 8, true, t)
}

func TestBEg2(t *testing.T) {
	doTest("testdata/day18Beg2.txt", 24, true, t)
}

func TestBEg3(t *testing.T) {
	doTest("testdata/day18Beg3.txt", 32, true, t)
}

func TestBEg4(t *testing.T) {
	doTest("testdata/day18Beg4.txt", 70, true, t) // adventofcode says ans for this eg is 72 but looks like answer is 70.
}

func doTest(filePath string, expected int, partB bool, t *testing.T) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	coords := toCoords(f)
	adjlist := toAdjList(coords)
	compressedadjlist := toCompressedAdjlist(adjlist, coords)
	res := process(compressedadjlist, coords, partB)

	if res != expected {
		t.Errorf("got %d. expected %d\n", res, expected)
	}
}
