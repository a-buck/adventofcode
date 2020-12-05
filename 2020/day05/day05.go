package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type seat struct {
	row int
	col int
}

var (
	inputFilePath = flag.String("input", "day05.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := run(file, *partB)

	fmt.Println(ans)
}

func run(r io.Reader, partB bool) int {
	scanner := bufio.NewScanner(r)

	maxID := 0

	ids := make([]int, 0)

	for scanner.Scan() {
		seat := decodeSeat(scanner.Text())
		if seat.id() > maxID {
			maxID = seat.id()
		}
		ids = append(ids, seat.id())
	}

	if partB {
		sort.IntSlice(ids).Sort()
		for i := 1; i < len(ids); i++ {
			diff := ids[i] - ids[i-1]
			if diff == 2 {
				return ids[i] - 1
			}
		}
		// seat not found
		return -1

	} else {
		// part A
		return maxID
	}
}

func decodeSeat(s string) seat {
	chars := []rune(s)

	row := binarySearch(chars[:7], 0, 127)
	col := binarySearch(chars[7:], 0, 7)

	return seat{row, col}
}

func binarySearch(chars []rune, lo, hi int) int {

	for i := 0; lo < hi; i++ {
		mid := lo + (hi-lo)/2
		c := chars[i]
		if c == 'F' || c == 'L' {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func (s seat) id() int {
	return s.row*8 + s.col
}
