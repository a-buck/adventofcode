package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day16.txt", "input file path")
	phases        = flag.Int("phases", 100, "number of phases")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	strSeq := strings.Split(string(inputBytes), "")

	repeats := 1
	offset := 0

	if *partB {
		repeats = 10000

		offset, err = strconv.Atoi(strings.Join(strSeq[:7], ""))
		if err != nil {
			log.Fatal(err)
		}

		if offset < (len(strSeq) * repeats / 2) {
			log.Fatalf("part b only works for offset >= n/2")
		}

		if offset > len(strSeq)*repeats-1 {
			log.Fatalf("offset: %d is larger than size of sequence: %d", offset, len(strSeq)*repeats)
		}

	}

	// extend input
	seq := make([]int, 0, len(strSeq)*repeats)
	for r := 0; r < repeats; r++ {
		for _, v := range strSeq {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			seq = append(seq, n)
		}
	}

	var res []int

	if *partB {
		res = sumBackwards(seq, *phases, offset)
	} else {
		// part A
		res = ftt(seq, *phases)
	}

	for i := offset; i < offset+8; i++ {
		fmt.Printf("%d", res[i])
	}
}

func sumBackwards(seq []int, phases, offset int) []int {
	for p := 0; p < phases; p++ { // phase
		for e := len(seq) - 2; e >= offset; e-- { // element
			seq[e] = abs(seq[e+1]+seq[e]) % 10
		}
	}
	return seq
}

func ftt(seq []int, phases int) []int {
	for p := 0; p < phases; p++ { // phase
		for e := range seq { // element
			total := 0
			for i, v := range seq {
				total += v * pattern(i, e)
			}
			seq[e] = abs(total) % 10 // unit digit only
		}
	}
	return seq
}

// pattern for ith value when
// considering element e
func pattern(i int, e int) int {
	pattern := []int{0, 1, 0, -1}

	index := ((i + 1) / (e + 1)) % 4

	return pattern[index]
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
