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
	var progInput = flag.Int("inputVal", 0, "program input")

	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	input := string(inputBytes)

	program := strings.Split(input, ",")

	outputs := run(tointSlice(program), *progInput)

	for _, v := range outputs {
		fmt.Println(v)
	}
}

func run(program []int, input int) []int {
	// program := tointSlice(programStrSlice)

	outputs := make([]int, 0)

	instrPtr := 0
loop:
	for instrPtr < len(program) {

		instr := program[instrPtr]
		instrStr := strconv.Itoa(instr)

		var opcode int
		firstParamImmediateMode := false
		secondParamImmediateMode := false
		thirdParamImmediateMode := false

		if len(instrStr) == 1 {
			opcode = toInt(instrStr)
		} else {
			// last 2 digits
			opcode = toInt(instrStr[len(instrStr)-2 : len(instrStr)])
		}

		if len(instrStr) >= 3 {
			firstParamImmediateMode = string(instrStr[len(instrStr)-3]) == "1"
		}
		if len(instrStr) >= 4 {
			secondParamImmediateMode = string(instrStr[len(instrStr)-4]) == "1"
		}
		if len(instrStr) >= 5 {
			thirdParamImmediateMode = string(instrStr[len(instrStr)-5]) == "1"
		}

		switch opcode {
		case 1:
			// add (3 = 1 + 2)
			firstParam := getParam(instrPtr+1, program, firstParamImmediateMode)
			secondParam := getParam(instrPtr+2, program, secondParamImmediateMode)
			result := firstParam + secondParam
			pIdx := instrPtr + 3
			updateProgramWithResult(pIdx, program, result, thirdParamImmediateMode)
			instrPtr += 4
		case 2:
			// multiply ( 3 = 1 * 2)
			firstParam := getParam(instrPtr+1, program, firstParamImmediateMode)
			secondParam := getParam(instrPtr+2, program, secondParamImmediateMode)
			result := firstParam * secondParam
			pIdx := instrPtr + 3
			updateProgramWithResult(pIdx, program, result, thirdParamImmediateMode)
			instrPtr += 4
		case 3:
			// input (1 = input)
			p := program[instrPtr+1]
			program[p] = input
			instrPtr += 2
		case 4:
			// output
			p := program[instrPtr+1]
			o := program[p]
			outputs = append(outputs, o)
			instrPtr += 2
		case 5:
			// jump-if-true
			firstParam := getParam(instrPtr+1, program, firstParamImmediateMode)
			secondParam := getParam(instrPtr+2, program, secondParamImmediateMode)

			if firstParam != 0 {
				instrPtr = secondParam
			} else {
				instrPtr += 3
			}
		case 6:
			// jump-if-false
			firstParam := getParam(instrPtr+1, program, firstParamImmediateMode)
			secondParam := getParam(instrPtr+2, program, secondParamImmediateMode)

			if firstParam == 0 {
				instrPtr = secondParam
			} else {
				instrPtr += 3
			}
		case 7:
			//less than
			firstParam := getParam(instrPtr+1, program, firstParamImmediateMode)
			secondParam := getParam(instrPtr+2, program, secondParamImmediateMode)

			pIdx := instrPtr + 3

			if firstParam < secondParam {
				updateProgramWithResult(pIdx, program, 1, thirdParamImmediateMode)
			} else {
				updateProgramWithResult(pIdx, program, 0, thirdParamImmediateMode)
			}

			instrPtr += 4
		case 8:
			firstParam := getParam(instrPtr+1, program, firstParamImmediateMode)
			secondParam := getParam(instrPtr+2, program, secondParamImmediateMode)

			pIdx := instrPtr + 3

			var result int
			if firstParam == secondParam {
				result = 1
			} else {
				result = 0
			}

			updateProgramWithResult(pIdx, program, result, thirdParamImmediateMode)

			instrPtr += 4
		case 99:
			//halt
			break loop
		default:
			log.Fatalf("Oops, unknown opcode %d", opcode)
		}
	}
	return outputs
}

func updateProgramWithResult(idx int, program []int, result int, isImmediateMode bool) {
	if !isImmediateMode {
		// position mode
		p := program[idx]
		program[p] = result
	} else {
		program[idx] = result
	}
}

func getParam(paramIndex int, program []int, isImmediateMode bool) int {
	p := program[paramIndex]
	if !isImmediateMode {
		// position mode
		return program[p]
	}
	return p
}

func tointSlice(s []string) []int {
	numbers := make([]int, len(s))

	for i, v := range s {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = val
	}

	return numbers
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
