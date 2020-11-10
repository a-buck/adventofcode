package intcode

import (
	"log"
	"strconv"
	"strings"
)

// Program ...
type Program []int

// ReadProgram ...
func ReadProgram(input []byte) Program {
	tokens := strings.Split(string(input), ",")
	program := tointSlice(tokens)
	return program
}

// Run intcode
// returns program state
func Run(program []int, inputs chan int, outputs chan int) []int {

	// extend slice...
	// todo find nicer way than just picking a value to extend by
	for i := 0; i < 10000; i++ {
		program = append(program, 0)
	}

	instrPtr := 0
	relativeBase := 0

loop:
	for instrPtr < len(program) {

		instr := program[instrPtr]

		opcode := instr % 100
		instr /= 100

		firstParamMode := instr % 10
		instr /= 10
		secondParamMode := instr % 10
		instr /= 10
		thirdParamMode := instr % 10

		switch opcode {
		case 1:
			// add
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			secondParam := getParam(instrPtr+2, program, secondParamMode, relativeBase)
			result := firstParam + secondParam
			updateProgramWithResult(instrPtr+3, program, result, thirdParamMode, relativeBase)
			instrPtr += 4
		case 2:
			// multiply
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			secondParam := getParam(instrPtr+2, program, secondParamMode, relativeBase)
			result := firstParam * secondParam
			updateProgramWithResult(instrPtr+3, program, result, thirdParamMode, relativeBase)
			instrPtr += 4
		case 3:
			// input
			inputVal := <-inputs
			updateProgramWithResult(instrPtr+1, program, inputVal, firstParamMode, relativeBase)
			instrPtr += 2
		case 4:
			// output
			o := getParam(instrPtr+1, program, firstParamMode, relativeBase) // new
			outputs <- o
			instrPtr += 2
		case 5:
			// jump-if-true
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			secondParam := getParam(instrPtr+2, program, secondParamMode, relativeBase)

			if firstParam != 0 {
				instrPtr = secondParam // todo maybe ptr should be 64 bit too?
			} else {
				instrPtr += 3
			}
		case 6:
			// jump-if-false
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			secondParam := getParam(instrPtr+2, program, secondParamMode, relativeBase)

			if firstParam == 0 {
				instrPtr = secondParam
			} else {
				instrPtr += 3
			}
		case 7:
			//less than
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			secondParam := getParam(instrPtr+2, program, secondParamMode, relativeBase)

			pIdx := instrPtr + 3

			var result int
			if firstParam < secondParam {
				result = 1
			} else {
				result = 0
			}

			updateProgramWithResult(pIdx, program, result, thirdParamMode, relativeBase)

			instrPtr += 4
		case 8:
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			secondParam := getParam(instrPtr+2, program, secondParamMode, relativeBase)

			pIdx := instrPtr + 3

			var result int
			if firstParam == secondParam {
				result = 1
			} else {
				result = 0
			}

			updateProgramWithResult(pIdx, program, result, thirdParamMode, relativeBase)

			instrPtr += 4

		case 9:
			// adjusts relative base
			firstParam := getParam(instrPtr+1, program, firstParamMode, relativeBase)
			relativeBase += firstParam
			instrPtr += 2
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

func updateProgramWithResult(idx int, program []int, result int, mode int, relativeBase int) {
	if mode == 0 {
		// position mode
		p := program[idx]
		program[p] = result
	} else if mode == 1 {
		// immediate mode
		program[idx] = result
	} else if mode == 2 {
		// relative mode
		p := program[idx]
		program[relativeBase+p] = result
	}
}

func getParam(paramIndex int, program []int, mode int, relativeBase int) int {
	p := program[paramIndex]
	var val int
	if mode == 0 {
		// position mode
		val = program[p]
	} else if mode == 1 {
		// immediate mode
		val = p
	} else if mode == 2 {
		// relative mode
		val = program[relativeBase+p]
	} else {
		panic("unknown param modemode")
	}

	return val
}

// Copy program
func (p Program) Copy() Program {
	cpy := make([]int, len(p))
	copy(cpy, p)
	return cpy
}
