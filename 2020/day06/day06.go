package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day06.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	groups := strings.Split(string(content), "\n\n")

	ans := run(groups, *partB)

	fmt.Println(ans)
}

func run(groups []string, partB bool) int {

	sum := 0

	for _, g := range groups {

		answers := make(map[rune]int) // answer to count of persons
		persons := strings.Split(g, "\n")

		for _, p := range persons {
			for _, a := range p {
				answers[a] = answers[a] + 1
			}
		}
		if partB {
			// everyone answered
			for _, v := range answers {
				if v == len(persons) {
					sum++
				}
			}

		} else {
			// part A
			// anyone answered
			sum += len(answers)
		}
	}
	return sum
}
