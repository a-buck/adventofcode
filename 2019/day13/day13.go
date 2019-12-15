package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

var (
	inputFilePath = flag.String("input", "day13.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	inputs := make(chan int, 1)
	outputs := make(chan int, 1)

	go intcode.Run(program, inputs, outputs)

	paramNumber := 0
	// var x int
	// var y int
	var tileID int

	numberBlocks := 0

	for i := range outputs {
		switch paramNumber {
		case 0:
			// x = i
		case 1:
			// y = i
		case 2:
			tileID = i
			if tileID == 2 {
				// block tile
				numberBlocks++
			}

		}
		paramNumber++
		if paramNumber == 3 {
			paramNumber = 0
		}
	}
	fmt.Println(numberBlocks)
}
