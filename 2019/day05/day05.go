package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

func main() {

	var inputFilePath = flag.String("input", "day5.txt", "input file path")
	var progInput = flag.Int("inputVal", 0, "program input")

	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	inputs := make(chan int, 1)
	outputs := make(chan int, 1)
	inputs <- *progInput

	go intcode.Run(program, inputs, outputs)

	for v := range outputs {
		fmt.Println(v)
	}
}
