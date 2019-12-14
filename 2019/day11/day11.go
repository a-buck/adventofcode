package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

var (
	inputFilePath = flag.String("input", "day11.txt", "input file path")
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

	camera <- 0 // starts off on black

	numberPanelsPaintedAtLeastOnce := 0

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

			if color == 1 {
				// white
				panels[currPosition] = 1
			} else {
				// black
				panels[currPosition] = 0
			}

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
			if currDirection == 0 {
				// go up
				currPosition = coord{currPosition.x, currPosition.y + 1}
			} else if currDirection == 1 {
				// go right
				currPosition = coord{currPosition.x + 1, currPosition.y}
			} else if currDirection == 2 {
				// go down
				currPosition = coord{currPosition.x, currPosition.y - 1}
			} else if currDirection == 3 {
				// go left
				currPosition = coord{currPosition.x - 1, currPosition.y}
			}

			// 4) send to camera color of new position
			camera <- panels[currPosition]

		}

		isFirstParam = !isFirstParam
	}

	fmt.Println(numberPanelsPaintedAtLeastOnce)

}
