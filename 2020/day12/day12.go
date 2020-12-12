package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day12.txt", "input file path")
)

type instruction struct {
	action rune
	value  int
}

type coord struct {
	row, col int
}

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	instructions := make([]instruction, len(lines))

	for i, v := range lines {

		action := rune(v[0])
		value, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}

		instructions[i] = instruction{action, value}

	}

	aAns, bAns := run(instructions)
	fmt.Println(aAns)
	fmt.Println(bAns)

}

// returns (part A, part B)
func run(instructions []instruction) (int, int) {

	directions := map[int]coord{
		0:   {row: -1}, // north
		90:  {col: 1},  // east
		180: {row: 1},  // south
		270: {col: -1}, // west
	}

	relativeWaypointPos := coord{row: -1, col: 10} // relative position from ship for part B.
	partBShipPos := coord{}                        // absolute position of ship for part B

	partAShipPos := coord{}
	partAShipBearing := 90

	for _, instr := range instructions {
		switch instr.action {
		case 'N':
			distance := scale(directions[0], instr.value)
			partAShipPos = add(partAShipPos, distance)
			relativeWaypointPos = add(relativeWaypointPos, distance)
		case 'S':
			distance := scale(directions[180], instr.value)
			partAShipPos = add(partAShipPos, distance)
			relativeWaypointPos = add(relativeWaypointPos, distance)
		case 'E':
			distance := scale(directions[90], instr.value)
			partAShipPos = add(partAShipPos, distance)
			relativeWaypointPos = add(relativeWaypointPos, distance)
		case 'W':
			distance := scale(directions[270], instr.value)
			partAShipPos = add(partAShipPos, distance)
			relativeWaypointPos = add(relativeWaypointPos, distance)

		case 'L':
			partAShipBearing = mod(partAShipBearing-instr.value, 360)
			relativeWaypointPos = rotateLeft(relativeWaypointPos, instr.value)
		case 'R':
			partAShipBearing = mod(partAShipBearing+instr.value, 360)
			relativeWaypointPos = rotateRight(relativeWaypointPos, instr.value)
		case 'F':
			partAShipPos = add(partAShipPos, scale(directions[partAShipBearing], instr.value))
			partBShipPos = add(partBShipPos, scale(relativeWaypointPos, instr.value))
		}
	}

	partA := abs(partAShipPos.col) + abs(partAShipPos.row)
	partB := abs(partBShipPos.col) + abs(partBShipPos.row)

	return partA, partB
}

func add(c1, c2 coord) coord {
	return coord{c1.row + c2.row, c1.col + c2.col}
}

func scale(c coord, factor int) coord {
	return coord{c.row * factor, c.col * factor}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func rotateLeft(c coord, theta int) coord {
	return rotateRight(c, 360-theta)
}

func rotateRight(c coord, degrees int) coord {

	for i := 0; i < degrees/90; i++ {
		c = coord{row: c.col, col: -c.row}
	}
	return c
}
