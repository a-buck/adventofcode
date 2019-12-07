package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/a-buck/adventofcode/2019/intcode"
)

func main() {
	var inputFilePath = flag.String("input", "day7.txt", "input file path")
	var partB = flag.Bool("partB", false, "enable part ")

	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(inputBytes)

	res := run(program, *partB)

	fmt.Println(res)
}

func run(program []int, partB bool) int {

	max := 0

	// Part A phases range
	start := 0
	end := 4

	if partB {
		start += 5
		end += 5
	}

	for a := start; a <= end; a++ {
		for b := start; b <= end; b++ {
			for c := start; c <= end; c++ {
				for d := start; d <= end; d++ {
					for e := start; e <= end; e++ {
						if allDifferent(a, b, c, d, e) {

							out := tryPhaseSequence(program, a, b, c, d, e)

							if out > max {
								max = out
							}
						}
					}
				}
			}
		}
	}
	return max
}

func allDifferent(vals ...int) bool {
	seen := make(map[int]bool, len(vals))

	for _, v := range vals {
		seen[v] = true
	}

	if len(seen) == len(vals) {
		return true
	}
	return false
}

func tryPhaseSequence(program []int, a int, b int, c int, d int, e int) int {

	var wg sync.WaitGroup
	wg.Add(5)

	eToA := make(chan int, 2)

	aToB := make(chan int, 1)
	bToC := make(chan int, 1)
	cToD := make(chan int, 1)
	dToE := make(chan int, 1)

	eToA <- a
	eToA <- 0

	aToB <- b
	bToC <- c
	cToD <- d
	dToE <- e

	runProgram := func(input chan int, output chan int) {
		defer wg.Done()
		runNewInstanceOfProgram(program, input, output)
	}

	go runProgram(eToA, aToB) // A
	go runProgram(aToB, bToC) // B
	go runProgram(bToC, cToD) // C
	go runProgram(cToD, dToE) // D
	go runProgram(dToE, eToA) // E

	wg.Wait()

	// get last output from channel written to by E
	var latestOutput int
	for i := range eToA {
		latestOutput = i
	}

	return latestOutput
}

func runNewInstanceOfProgram(program []int, input chan int, output chan int) {
	programcopy := make([]int, len(program))
	copy(programcopy, program)
	intcode.Run(programcopy, input, output)
}
