package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

var (
	inputFilePath = flag.String("input", "day19.txt", "input file path")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	count := 0
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			// intcode halts after each cycle
			// so we create a new intcode computer each time
			inputs := make(chan int)
			outputs := make(chan int)
			go intcode.Run(program.Copy(), inputs, outputs)
			inputs <- i // x
			inputs <- j // y

			if <-outputs == 1 { // being pulled
				count++
			}
		}
	}
	fmt.Println(count)
}
