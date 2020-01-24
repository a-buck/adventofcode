package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

type coord struct {
	x int
	y int
}

var (
	inputFilePath = flag.String("input", "day19.txt", "input file path")
	squareSize    = flag.Int("squareSize", 100, "Size of square to fit in tractor beam")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)
	coords := make(map[coord]rune, 0)

	count := 0
	// for every row
	for y := 0; y < 50; y++ {
		// for every col
		for x := 0; x < 50; x++ {
			if isTractorBeam(coord{x, y}, program) {
				count++
				coords[coord{x, y}] = '#'
			}
		}
	}

	printCoords(coords, 30)
	fmt.Printf("Part A: %d\n", count)

	down := coord{y: 1}
	right := coord{x: 1}

	curr := coord{x: 3, y: 4} // starting point. todo: avoid hardcoding this.

outer:
	for {
		// go down
		curr = curr.add(down)
		// go right until hit a tractor beam
		for {
			// if we find part of the beam, we are in the bottoms left corner of the 100 x 100 square
			if isTractorBeam(curr, program) {
				diff := *squareSize - 1
				bottomRight := curr.add(coord{x: diff})
				topLeft := curr.add(coord{y: -diff})
				topRight := curr.add(coord{x: diff, y: -diff})
				// check if all 4 corners of square are in tractor beams
				if curr.y-diff >= 0 &&
					isTractorBeam(bottomRight, program) &&
					isTractorBeam(topLeft, program) &&
					isTractorBeam(topRight, program) {
					fmt.Printf("Part B: %d\n", topLeft.x*10000+topLeft.y)
					break outer
				}
				break // stop going rightf
			}
			curr = curr.add(right)
		}
	}
}

func isTractorBeam(c coord, program intcode.Program) bool {
	program = program.Copy()
	inputs := make(chan int)
	outputs := make(chan int)
	go intcode.Run(program, inputs, outputs)
	inputs <- c.x
	inputs <- c.y
	return 1 == <-outputs
}

func (c coord) add(other coord) coord {
	return coord{c.x + other.x, c.y + other.y}
}

func printCoords(coords map[coord]rune, size int) {
	for y := 0; y < size; y++ {
		fmt.Printf("%02d", y)
		for x := 0; x < size; x++ {
			c := coord{x, y}
			v, ok := coords[c]
			if ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
