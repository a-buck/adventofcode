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

	ans := run(file, *partB)

	fmt.Printf("%d\n", ans)

}

func run(r io.Reader, partB bool) int {
	const target = 2020

	vals := make(map[int]int)

	scanner := bufio.NewScanner(r)

	// populate vals with value being amount needed to reach target
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		vals[val] = target - val
	}

	for k, v := range vals {
		_, ok := vals[v]

		if ok {
			return k * v
		}
	}

	return 0
}
