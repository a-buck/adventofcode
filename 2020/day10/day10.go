package main

import (
	"flag"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day10.txt", "input file path")
)

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	numbers := make([]int, len(lines))

	for i, v := range lines {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = n
	}

	aAns := partA(numbers)
	println(aAns)

	bAns := partB(numbers)
	println(bAns)
}

func partA(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	sort.Ints(numbers)

	diffs := make(map[int]int) // map of diff to count

	diffs[numbers[0]]++ // diff of outlet (0) to first device
	diffs[3]++          // diff from last adapter to device is always 3

	for i := 1; i < len(numbers); i++ {
		curr := numbers[i]
		prev := numbers[i-1]

		diffs[curr-prev]++
	}

	return diffs[1] * diffs[3]
}

func partB(numbers []int) int {
	sort.Ints(numbers)
	adjlist := buildAdjlist(numbers)
	reversed := reverseAdjlist(adjlist)

	childPathsCount := make(map[int]int)

	for i := len(numbers) - 1; i >= 0; i-- {
		child := numbers[i]
		parents := reversed[child]

		for _, p := range parents {

			inc := 0

			if childPathsCount[child] == 0 {
				inc = 1
			}

			inc += childPathsCount[child]

			childPathsCount[p] += inc
		}

	}

	return childPathsCount[0]
}

func buildAdjlist(numbers []int) map[int][]int {

	numbersmap := make(map[int]bool)

	max := 0

	for _, n := range numbers {
		numbersmap[n] = true
		if n > max {
			max = n
		}
	}

	adjlist := make(map[int][]int)

	for _, v := range numbers {
		for i := 1; i <= 3; i++ {
			if _, ok := numbersmap[v+i]; ok {
				adjlist[v] = append(adjlist[v], v+i)
			}
		}
	}

	// add neighbours from 0
	for i := 1; i <= 3; i++ {
		if _, ok := numbersmap[i]; ok {
			adjlist[0] = append(adjlist[0], i)
		}
	}

	// add neighbours from max to device
	adjlist[max] = append(adjlist[max], max+3)

	return adjlist
}

func reverseAdjlist(adjlist map[int][]int) map[int][]int {
	reversed := make(map[int][]int)

	for k, v := range adjlist {
		for _, v2 := range v {
			reversed[v2] = append(reversed[v2], k)
		}
	}

	return reversed
}
