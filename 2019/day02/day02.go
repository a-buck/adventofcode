package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

func main() {
	var inputFilePath = flag.String("input", "day1.txt", "input file path")
	var partB = flag.Bool("partB", false, "enable part b")

	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	if *partB {
		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				programcopy := make([]int, len(program))
				copy(programcopy, program)
				programcopy[1] = noun
				programcopy[2] = verb
				input := make(chan int, 1)
				output := make(chan int, 1)
				prog := intcode.Run(programcopy, input, output)
				if prog[0] == 19690720 {
					answer := 100*noun + verb
					fmt.Println(answer)
				}
			}

		}
	} else {
		// part A
		input := make(chan int, 1)
		input <- 0
		output := make(chan int, 1)
		program[1] = 12
		program[2] = 2
		prog := intcode.Run(program, input, output)
		fmt.Println(prog[0])
	}

}
