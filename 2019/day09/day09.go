package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

var (
	inputFilePath = flag.String("input", "day09.txt", "input file path")
	progInput     = flag.Int("progInput", 1, "program input")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	inputs := make(chan int, 1)
	inputs <- int(*progInput)
	outputChannel := make(chan int, 1)

	go intcode.Run(program, inputs, outputChannel)

	outputs := make([]int, 0)
	for v := range outputChannel {
		outputs = append(outputs, v)
	}

	for _, v := range outputs {
		fmt.Printf("%d,", v)
	}
}
