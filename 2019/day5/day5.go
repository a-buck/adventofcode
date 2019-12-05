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

	var inputFilePath = flag.String("input", "day5.txt", "input file path")
	var partB = flag.Bool("partB", false, "enable part b") // todo currently this represents day 2 part b

	var inputVal = flag.Int("inputVal", 0, "program input")

	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	input := string(inputBytes)

	splitInput := strings.Split(input, ",")

	answer, err := runProgram(splitInput, *partB, *inputVal) // todo: is day 5 using part b of day2 ?

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}

func runProgram(program []string, partbnotusedyet bool, input int) (int, error) {
	numbers := make([]int, len(program))

	for i, v := range program {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = val
	}

	return evaluate(numbers, input)
}

func evaluate(numbers []int, input int) (int, error) {
	i := 0
loop:
	for i < len(numbers) {

		switch opcode := numbers[i]; opcode {
		case 1:
			// add (3 = 1 + 2)
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
			i += 4
		case 2:
			// multiply ( 3 = 1 + 2)
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
			i += 4
		case 3:
			// input (1 = input)
			numbers[numbers[i+1]] = input
			i += 2
		case 4:
			// output
			return numbers[numbers[i+1]], nil
		case 99:
			//halt
			break loop
		}
	}
	return 0, fmt.Errorf("no output")
}
