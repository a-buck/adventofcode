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
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	strSeq := strings.Split(string(inputBytes), "")

	intSeq := make([]int, len(strSeq))

	for i, v := range strSeq {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		intSeq[i] = n
	}

	res := ftt(intSeq, *phases)

	resStr := make([]string, len(res))
	for i, v := range res {
		resStr[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(resStr[:8], ""))
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
