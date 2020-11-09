package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/a-buck/adventofcode/2019/intcode"
)

type direction int

type coord struct {
	x int
	y int
}

var (
	inputFilePath = flag.String("input", "day17.txt", "input file path")
	videoFeed     = flag.Bool("videoFeed", false, "enable continuous video feed")
)

var (
	deltas          = [4]coord{{y: -1}, {x: 1}, {y: 1}, {x: -1}} // up, right, down, left
	robotDirections = map[rune]direction{
		'^': 0,
		'>': 1,
		'v': 2,
		'<': 3,
	}
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	partAProgram := intcode.ReadProgram(inputBytes)
	partBProgram := partAProgram.Copy()

	partAInputs := make(chan int)
	partAOutputs := make(chan int, 1)

	go intcode.Run(partAProgram, partAInputs, partAOutputs)

	row, col := 0, 0

	coords := make(map[coord]rune)

	for o := range partAOutputs {
		fmt.Printf("%c", rune(o))
		if rune(o) == '\n' {
			row++
			col = 0
		} else {
			if rune(o) != '.' {
				coords[coord{col, row}] = rune(o)
			}
			col++
		}
	}

	var robot coord
	var robotDirection direction

	alignmentParameterSum := 0
	for k, v := range coords {

		surroundedByScaffolding := true
		for _, n := range k.neighboours() {
			if coords[n] != '#' {
				surroundedByScaffolding = false
			}
		}

		if v == '#' && surroundedByScaffolding { // intersection
			alignmentParameterSum += k.x * k.y
		}

		if d, ok := robotDirections[v]; ok {
			robotDirection = d
			robot = k
		}
	}
	fmt.Printf("Part A: %d\n", alignmentParameterSum)

	instrs := make([]string, 0)
	countAcc := 0

outer:
	for {
		next := robot.addDirection(robotDirection)

		if v, ok := coords[next]; v == '#' {
			countAcc++
			robot = next
		} else if !ok {
			if countAcc > 0 {
				instrs = append(instrs, strconv.Itoa(countAcc))
			}
			countAcc = 0
			neighbours := robot.neighboours()
			for i := 1; i < 4; i++ {

				dir := direction((int(robotDirection) + i) % 4)
				n := neighbours[dir]
				if coords[n] != '#' {
					continue
				}

				if i != 2 { // not opposite

					var instr string
					if i == 1 {
						instr = "R"
					} else if i == 3 {
						instr = "L"
					} else {
						log.Fatalf("unexpected direction %d", i)
					}

					robotDirection = dir
					instrs = append(instrs, instr)
					continue outer
				}
			}
			break // no neighbours - must be at end.
		}
	}

	fmt.Printf("Instructions: %+v\n", instrs)

	// part b

	partBInputs := make(chan int, 1)
	partBOutputs := make(chan int, 1)

	partBProgram[0] = 2 // enable interactive mode

	go intcode.Run(partBProgram, partBInputs, partBOutputs)

	done := make(chan bool)
	go func(done chan bool) {
		var prev int
		for o := range partBOutputs {
			prev = o
			fmt.Printf("%c", rune(o))
		}
		fmt.Printf("\nLast output value: %d\n", prev)
		done <- true
	}(done)

	// todo: automatically  determining what main routine and functions A,B,C should be.
	reader := bufio.NewReader(os.Stdin)
	mainRoutine, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	sendToChannel(mainRoutine, partBInputs)
	a, _ := reader.ReadString('\n')
	sendToChannel(a, partBInputs)
	b, _ := reader.ReadString('\n')
	sendToChannel(b, partBInputs)
	c, _ := reader.ReadString('\n')

	sendToChannel(c, partBInputs)

	if *videoFeed {
		sendToChannel("y\n", partBInputs)
	} else {
		sendToChannel("n\n", partBInputs)
	}

	<-done
}

func sendToChannel(s string, c chan int) {
	rs := []rune(s)

	for _, r := range rs {
		c <- int(r)
	}

}

func (c coord) addCoord(other coord) coord {
	return coord{c.x + other.x, c.y + other.y}
}

func (c coord) addDirection(d direction) coord {
	switch d {
	case 0:
		c.y--
	case 1:
		c.x++
	case 2:
		c.y++
	case 3:
		c.x--
	}
	return c
}

func (c coord) neighboours() []coord {

	coords := make([]coord, 0, 4)

	for _, d := range deltas {
		added := c.addCoord(d)
		coords = append(coords, added)
	}
	return coords
}
