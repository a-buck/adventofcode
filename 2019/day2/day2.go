package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var inputFilePath = flag.String("input", "day1.txt", "input file path")
	var partB = flag.Bool("partB", false, "enable part b")

	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	input := string(inputBytes)

	splitInput := strings.Split(input, ",")

	numbers := make([]int, len(splitInput))

	for i, v := range splitInput {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = val
	}

	if *partB {
		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				cpy := make([]int, len(numbers))
				copy(cpy, numbers)
				res := evaluate(cpy, noun, verb)
				if res == 19690720 {
					answer := 100*noun + verb
					fmt.Println(answer)
				}
			}

		}
	} else {
		// part A
		res := evaluate(numbers, 12, 2)
		fmt.Println(res)
	}

}

func evaluate(numbers []int, noun int, verb int) int {

	numbers[1] = noun
	numbers[2] = verb

loop:
	for i := 0; i < len(numbers); i += 4 {

		switch n := numbers[i]; n {
		case 1:
			// add
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
		case 2:
			// multiply
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
		case 99:
			//halt
			break loop
		}
	}
	return numbers[0]
}
