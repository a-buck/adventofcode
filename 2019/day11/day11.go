package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"github.com/a-buck/adventofcode/2019/intcode"
)

var (
	inputFilePath = flag.String("input", "day11.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type coord struct {
	x int
	y int
}

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	camera := make(chan int, 1)
	output := make(chan int, 1)

	go intcode.Run(program, camera, output)

	isFirstParam := true
	var color int
	var direction int

	currDirection := 0
	var currPosition coord
	panels := make(map[coord]int) // color of panels

	if *partB {
		// start on white panel
		panels[currPosition] = 1
	}

	camera <- panels[currPosition]

	numberPanelsPaintedAtLeastOnce := 0

	directions := map[int]coord{
		// up
		0: coord{x: 0, y: 1},
		// right
		1: coord{x: 1, y: 0},
		// down
		2: coord{x: 0, y: -1},
		// left
		3: coord{x: -1, y: 0},
	}

	for i := range output {
		if isFirstParam {
			color = i
		} else {
			direction = i

			// 1) paint current panel color
			_, ok := panels[currPosition]
			if !ok {
				// painting a new panel
				numberPanelsPaintedAtLeastOnce++
			}
			panels[currPosition] = color

			// 2) change direction
			if direction == 0 {
				// turn left
				currDirection--
				if currDirection == -1 {
					currDirection = 3
				}
			} else {
				// turn right
				currDirection++
				if currDirection == 4 {
					currDirection = 0
				}
			}

			// 3) move forward 1
			deltaPos := directions[currDirection]
			currPosition = currPosition.add(deltaPos)

			// 4) send to camera color of new position
			camera <- panels[currPosition]
		}

		isFirstParam = !isFirstParam
	}

	if *partB {
		print(panels)
	} else {
		// part A
		fmt.Println(numberPanelsPaintedAtLeastOnce)
	}

}

func print(panels map[coord]int) {

	maxX := math.MinInt32
	maxY := math.MinInt32

	minX := math.MaxInt32
	minY := math.MaxInt32

	for p := range panels {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}

	for row := maxY; row >= minY; row-- {
		for col := minX; col < maxX; col++ {
			color := panels[coord{x: col, y: row}]
			if color == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}

func (c coord) add(other coord) coord {
	return coord{x: c.x + other.x, y: c.y + other.y}
}
