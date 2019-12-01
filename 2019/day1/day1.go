package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"fmt"
)

func main() {

	var inputFilePath = flag.String("input", "day1.txt", "input file path")
	var _ = flag.Bool("partB", false, "enable part b")

	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fuel := 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		f := mass/3 - 2 // divide by 3, round down, subtract 2
		fuel += f
	}

	fmt.Println(fuel)
}
