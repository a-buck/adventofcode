package intcode

import (
	"log"
	"strconv"
	"strings"
)

// Read program
func ReadProgram(input []byte) []int {
	tokens := strings.Split(string(input), ",")
	program := tointSlice(tokens)
	return program
}

// Run intcode
// returns program state
func Run(program []int, inputs chan int, outputs chan int) []int {

	instrPtr := 0

loop:
	for instrPtr < len(program) {

		instr := program[instrPtr]

		opcode := instr % 100
		instr /= 100

		firstParamImmediateMode := instr%10 == 1
		instr /= 10
		secondParamImmediateMode := instr%10 == 1
		instr /= 10
		thirdParamImmediateMode := instr%10 == 1

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
			program[p] = <-inputs
			instrPtr += 2
		case 4:
			// output
			p := program[instrPtr+1]
			o := program[p]
			outputs <- o
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

			var result int
			if firstParam < secondParam {
				result = 1
			} else {
				result = 0
			}

			updateProgramWithResult(pIdx, program, result, thirdParamImmediateMode)

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
			close(outputs)
			break loop
		default:
			log.Fatalf("Oops, unknown opcode %d", opcode)
		}
	}
	return program
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
