package main

import (
	"log"
	"os"
	"testing"
)

func TestGetPortal(t *testing.T) {
	// above outer
	doGetPortalTest("testdata/day20b.txt", 13, 2, portal{"ZZ", true}, t)
	doGetPortalTest("testdata/day20b.txt", 15, 2, portal{"LP", true}, t)
	doGetPortalTest("testdata/day20b.txt", 17, 2, portal{"XQ", true}, t)
	doGetPortalTest("testdata/day20b.txt", 19, 2, portal{"WB", true}, t)
	doGetPortalTest("testdata/day20b.txt", 27, 2, portal{"CK", true}, t)

	// below outer
	doGetPortalTest("testdata/day20b.txt", 15, 34, portal{"AA", true}, t)
	doGetPortalTest("testdata/day20b.txt", 17, 34, portal{"OA", true}, t)
	doGetPortalTest("testdata/day20b.txt", 19, 34, portal{"FD", true}, t)
	doGetPortalTest("testdata/day20b.txt", 23, 34, portal{"NM", true}, t)

	// left outer
	doGetPortalTest("testdata/day20b.txt", 2, 15, portal{"CJ", true}, t)
	doGetPortalTest("testdata/day20b.txt", 2, 21, portal{"XF", true}, t)
	doGetPortalTest("testdata/day20b.txt", 2, 25, portal{"RE", true}, t)

	// right outer
	doGetPortalTest("testdata/day20b.txt", 42, 13, portal{"ZH", true}, t)
	doGetPortalTest("testdata/day20b.txt", 42, 17, portal{"IC", true}, t)
	doGetPortalTest("testdata/day20b.txt", 42, 25, portal{"RF", true}, t)

	// below inner
	doGetPortalTest("testdata/day20b.txt", 13, 8, portal{"FD", false}, t)
	doGetPortalTest("testdata/day20b.txt", 21, 8, portal{"RE", false}, t)
	doGetPortalTest("testdata/day20b.txt", 23, 8, portal{"IC", false}, t)
	doGetPortalTest("testdata/day20b.txt", 31, 8, portal{"ZH", false}, t)

	// above inner
	doGetPortalTest("testdata/day20b.txt", 17, 28, portal{"XF", false}, t)
	doGetPortalTest("testdata/day20b.txt", 21, 28, portal{"XQ", false}, t)
	doGetPortalTest("testdata/day20b.txt", 29, 28, portal{"LP", false}, t)

	// right inner
	doGetPortalTest("testdata/day20b.txt", 8, 13, portal{"OA", false}, t)
	doGetPortalTest("testdata/day20b.txt", 8, 17, portal{"CK", false}, t)
	doGetPortalTest("testdata/day20b.txt", 8, 23, portal{"CJ", false}, t)

	// left inner
	doGetPortalTest("testdata/day20b.txt", 36, 13, portal{"WB", false}, t)
	doGetPortalTest("testdata/day20b.txt", 36, 21, portal{"RF", false}, t)
	doGetPortalTest("testdata/day20b.txt", 36, 23, portal{"NM", false}, t)
}

func doGetPortalTest(input string, x, y int, expected portal, t *testing.T) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	grid := readInput(f)

	p := getPortal(grid, x, y)
	if p == nil {
		t.Error("portal is nil")
		return
	}
	if *p != expected {
		t.Errorf("p= %+v != expected=%+v", p, expected)
	}
}

func TestBPartBE2e(t *testing.T) {
	input := "testdata/day20b.txt"
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parsedInput := readInput(f)

	ans := process(parsedInput, true)

	if ans != 396 {
		t.Errorf("expected 396, got %d", ans)
	}
}

func TestPartAE2e(t *testing.T) {
	input := "testdata/day20a.txt"
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parsedInput := readInput(f)

	ans := process(parsedInput, false)

	if ans != 58 {
		t.Errorf("expected 58, got %d", ans)
	}
}
