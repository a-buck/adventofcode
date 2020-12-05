package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
		seatID := getSeatID(scanner.Text())
		if seatID > maxID {
			maxID = seatID
		}
		ids = append(ids, seatID)
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

func getSeatID(s string) int {
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "L", "0")
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "R", "1")

	val, err := strconv.ParseInt(s, 2, 0)

	if err != nil {
		log.Fatal(err)
	}
	return int(val)
}
