package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day08.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type instruction struct {
	// acc, jmp, nop
	operation string
	// signed number
	argument int
}

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	program := strings.Split(string(content), "\n")

	instructions := make([]instruction, 0, len(program))

	for _, instr := range program {
		parts := strings.Split(instr, " ")
		operation := parts[0]
		argument, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, instruction{operation: operation, argument: argument})
	}

	if *partB {
		result, err := doPartB(instructions)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)

	} else {
		// part A
		result, correctTermination := run(instructions)
		if correctTermination {
			fmt.Printf("expected termination due to repeating instruction for part A. Result: %d", result)
		} else {
			fmt.Println(result)
		}
	}
}

// find and update corrupted instruction
// returns accumulator for program after corruption is fixed
// returns error if corruption not found
func doPartB(instructions []instruction) (int, error) {
	for i := 0; i < len(instructions); i++ {
		instructionscpy := make([]instruction, len(instructions))
		copy(instructionscpy, instructions)

		instr := instructionscpy[i]

		if instr.operation == "jmp" {
			instructionscpy[i].operation = "nop"
		} else if instr.operation == "nop" {
			instructionscpy[i].operation = "jmp"
		} else {
			continue
		}
		result, correctTermination := run(instructionscpy)
		if correctTermination {
			return result, nil
		}
	}

	return 0, errors.New("corrupted instruction not found")
}

// returns (acc, correctTermination)
// where correctTermination is the case where instruction pointer
// reaches past end. False if there is a repeat of an instruction
func run(instructions []instruction) (int, bool) {

	seenInstrs := make(map[int]bool)

	acc := 0

	i := 0

	for i < len(instructions) {
		instr := instructions[i]

		if _, ok := seenInstrs[i]; ok {
			return acc, false
		}

		seenInstrs[i] = true

		switch instr.operation {
		case "acc":
			acc += instr.argument
			i++
		case "jmp":
			i += instr.argument
		case "nop":
			i++
		}
	}
	return acc, true

}
