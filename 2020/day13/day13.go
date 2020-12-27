package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"github.com/a-buck/adventofcode/2020/utils"
)

var (
	inputFilePath = flag.String("input", "day13.txt", "input file path")
)

type bus struct {
	id, offset int
}

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	earliestTs := utils.ToInt(lines[0])

	parts := strings.Split(lines[1], ",")
	buses := make([]bus, 0)
	for i, v := range parts {
		if v == "x" {
			continue
		}

		buses = append(buses, bus{id: utils.ToInt(v), offset: i})
	}

	ans := partA(buses, earliestTs)
	fmt.Println(ans)

	ansB := partB(buses)
	fmt.Println(ansB)

}

func partA(buses []bus, ts int) int {
	soonestTS := math.MaxInt32
	soonestID := 0
	for _, b := range buses {
		nextTs := ts / b.id * b.id
		if ts%b.id != 0 {
			nextTs += b.id
		}
		if nextTs < soonestTS {
			soonestTS = nextTs
			soonestID = b.id
		}
	}
	waitTime := soonestTS - ts
	return waitTime * soonestID
}

func partB(buses []bus) int {
	// chinese remainder theorem https://www.youtube.com/watch?v=zIFehsBHB8o
	// https://docs.google.com/spreadsheets/d/1GW-wnDxz9UF6woYPe4x0zoMwXWV1o4smt7MYgBSLpKw/edit#gid=0
	N := 1
	for _, b := range buses {
		N *= b.id
	}

	sum := 0
	for _, b := range buses {
		bi := b.id - b.offset
		Ni := N / b.id
		xi := findInverse(Ni, b.id)
		product := bi * Ni * xi
		sum += product
	}

	r := sum % N
	return r
}

func findInverse(Ni, ni int) int {
	// Ni . x ≡ 1 (mod ni)

	for x := 1; ; x++ {
		left := Ni * x

		if left%ni == 1%ni {
			return x
		}
	}
}
