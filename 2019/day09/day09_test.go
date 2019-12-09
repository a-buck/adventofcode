package main

import (
	"strconv"
	"testing"

	"github.com/a-buck/adventofcode/2019/intcode"
)

func TestPartAEchoProgramEg1(t *testing.T) {
	prog := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)

	go intcode.Run(prog, inputChannel, outputChannel)

	outputs := make([]int, 0)
	for i := range outputChannel {
		outputs = append(outputs, i)
	}

	// Check output is same as input
	for i, v := range prog {
		if outputs[i] != v {
			t.Errorf("got output[%d]=%d, expected prog[%d]=%d", i, outputs[i], i, v)
		}
	}

}

func TestPartAEg2(t *testing.T) {
	prog := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	inputs := make(chan int, 0)
	outputs := make(chan int, 0)
	go intcode.Run(prog, inputs, outputs)

	var last int
	for i := range outputs {
		last = i
	}

	if s := strconv.Itoa(last); len(s) != 16 {
		t.Errorf("got result of len %s, expected len 16", s)
	}
}

func TestPartAOutputLargeNumberEg3(t *testing.T) {
	largeVal := 1125899906842624
	prog := []int{104, largeVal, 99}
	inputs := make(chan int, 1)
	outputs := make(chan int, 1)

	go intcode.Run(prog, inputs, outputs)

	var last int
	for i := range outputs {
		last = i
	}

	if last != largeVal {
		t.Errorf("got %d, expected %d", last, largeVal)
	}
}
