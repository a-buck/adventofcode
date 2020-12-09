package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	inputFilePath  = flag.String("input", "day08.txt", "input file path")
	preambleLength = flag.Int("preamble", 25, "preamble length")
)

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(content), "\n")

	numbers := make([]int, len(rows))
	for i, row := range rows {
		v, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = v
	}

	partAResult := doPartA(numbers, *preambleLength)
	fmt.Println(partAResult)

	partBResult := doPartB(numbers, partAResult)
	fmt.Println(partBResult)
}

func doPartA(numbers []int, preambleLength int) int {
	if preambleLength > len(numbers) {
		log.Fatalf("preamble length: %d is longer than input length: %d", preambleLength, len(numbers))
	}

	window := make(map[int]bool)

	for i := 0; i < preambleLength; i++ {
		val := numbers[i]
		window[val] = true
	}

	for hi := preambleLength; hi < len(numbers); hi++ {
		lo := hi - preambleLength

		target := numbers[hi]

		found := false
		for _, v := range numbers[lo:hi] {
			other := target - v
			if window[other] && v != other {
				found = true
				break
			}
		}

		if !found {
			return target
		}

		// shift window along
		delete(window, numbers[lo])
		window[numbers[hi]] = true
	}

	return -1
}

func doPartB(numbers []int, target int) int {
	lo := 0
	hi := 0

	sum := numbers[0]

	for hi < len(numbers) {
		if sum < target {
			hi++
			sum += numbers[hi]

		} else if sum > target {
			sum -= numbers[lo]
			lo++
		}

		if sum == target {
			break
		}
	}

	min := math.MaxInt32
	max := math.MinInt32

	for i := lo; i <= hi; i++ {
		v := numbers[i]
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min + max
}
