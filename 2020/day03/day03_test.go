package main

import (
	"log"
	"os"
	"testing"
)

func TestAEg(t *testing.T) {
	f, err := os.Open("testdata/day03eg.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ans := run(f, false)

	if ans != 7 {
		t.Errorf("got %d, expected %d", ans, 7)
	}
}
