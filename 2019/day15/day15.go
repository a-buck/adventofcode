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

type node struct {
	coord coord
	depth int
}

type movement int

var (
	inputFilePath = flag.String("input", "day15.txt", "input file path")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	inputs := make(chan int)
	outputs := make(chan int)

	go intcode.Run(program, inputs, outputs)

	coords := make(map[coord]int)

	root := coord{}
	coords[root] = 1 // root is empty space

	oxygen := &node{}

	err = findOxygen(root, coords, inputs, outputs, 0, oxygen)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part A: %d\n", oxygen.depth)

	depth := findDepth(oxygen.coord, coords)

	fmt.Printf("Part B: %d\n", depth)
}

// dfs to find oxygen coordinates and depth
func findOxygen(c coord, coords map[coord]int, inputs, outputs chan int, depth int, oxygenNode *node) error {

	for i, n := range c.neighbours() {
		// if haven't visited neighbour
		if _, ok := coords[n]; !ok {
			cmd := i + 1 // convert from 0 to 1 based index

			inputs <- cmd
			status := <-outputs

			coords[n] = status

			if status == 2 { // oxygen
				*oxygenNode = node{n, depth + 1}
			}

			if status != 0 { // move if didn't hit wall
				findOxygen(n, coords, inputs, outputs, depth+1, oxygenNode)

				// backtrack
				inputs <- movement(cmd).opposite()
				backtrackOut := <-outputs

				if backtrackOut == 0 {
					return fmt.Errorf("Hit wall while backtracking")
				}
			}
		}

	}
	return nil
}

// bfs to find number of levels
func findDepth(root coord, coords map[coord]int) int {
	visited := make(map[coord]bool, 0)

	coordQ := make([]coord, 0)
	coordQ = append(coordQ, root)

	depthQ := []int{0}

	var depth int
	for len(coordQ) > 0 {
		curr := coordQ[0]
		coordQ = coordQ[1:]
		depth = depthQ[0]
		depthQ = depthQ[1:]

		for _, n := range curr.neighbours() {

			status, ok := coords[n]
			_, visited := visited[n]

			if status != 0 && ok && !visited {
				coordQ = append(coordQ, n)
				depthQ = append(depthQ, depth+1)
			}

		}
		visited[curr] = true
	}

	return depth
}

func (m movement) opposite() int {
	var opposite int
	switch m {
	case 1:
		opposite = 2
	case 2:
		opposite = 1
	case 3:
		opposite = 4
	case 4:
		opposite = 3
	}
	return opposite
}

// get neighbouring coordinates
// in order north, south, west, east
func (c coord) neighbours() []coord {
	return []coord{
		coord{c.x, c.y + 1},
		coord{c.x, c.y - 1},
		coord{c.x - 1, c.y},
		coord{c.x + 1, c.y},
	}
}
