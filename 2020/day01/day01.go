package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	inputFilePath = flag.String("input", "day01.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vals := readVals(file)
	ans := run(vals, *partB)

	fmt.Printf("%d\n", ans)
}

func readVals(r io.Reader) []int {
	vals := make([]int, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		vals = append(vals, val)
	}
	return vals

}

func run(vals []int, partB bool) int {
	const target = 2020

	valsToIndex := make(map[int]int)
	for i, v := range vals {
		valsToIndex[v] = i
	}

	if partB {
		for i, v1 := range vals {
			for j, v2 := range vals[i+1:] {
				v3 := target - v1 - v2
				k, ok := valsToIndex[v3]
				if ok && i != k && j != k {
					return v1 * v2 * v3
				}
			}
		}

	} else {
		// part A

		for i, v1 := range vals {
			v2 := target - v1
			j, ok := valsToIndex[v2]
			if ok && i != j {
				return v1 * v2
			}
		}
	}

	return 0
}
