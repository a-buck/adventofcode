package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type grid [][]rune

type coord struct {
	row, col int
}

var (
	inputFilePath = flag.String("input", "day11.txt", "input file path")
	partB         = flag.Bool("partB", false, "part B")
)

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	g := toGrid(content)
	hasChanged := true
	for hasChanged {
		g, hasChanged = g.next(*partB)
	}
	fmt.Println(g.occupiedCount())
}

func toGrid(content []byte) grid {
	lines := strings.Split(string(content), "\n")

	g := make(grid, len(lines))

	for i, l := range lines {
		g[i] = []rune(l)
	}

	return g
}

func directions() []coord {
	neighbours := make([]coord, 0, 8)
	for _, row := range []int{-1, 0, 1} {
		for _, cols := range []int{-1, 0, 1} {
			if cols == 0 && row == 0 {
				continue
			}
			neighbours = append(neighbours, coord{row, cols})
		}
	}
	return neighbours

}

func add(c1, c2 coord) coord {
	return coord{c1.row + c2.row, c1.col + c2.col}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (c coord) neighbours() []coord {
	dirs := directions()

	neighbours := make([]coord, 0, len(dirs))

	for _, d := range dirs {
		neighbours = append(neighbours, add(c, d))
	}
	return neighbours
}

// returns new grid and whether it is different to g
func (g grid) next(partB bool) (grid, bool) {
	to := g.copy()

	changesTotal := 0

	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			seatChange := g.seatStateChange(coord{row, col}, partB)

			if seatChange == 1 {
				to[row][col] = '#'
			} else if seatChange == -1 {
				to[row][col] = 'L'
			}
			changesTotal += abs(seatChange)
		}
	}

	return to, changesTotal > 0

}

// 1 if seat becomes occupied
// 0 if no change
// -1 if seat becomes empty
func (g grid) seatStateChange(c coord, partB bool) int {
	seat := g.get(c)

	if seat == '.' {
		// floor
		return 0
	}

	dirs := directions()

	var neighbourFn func(c, dir coord) rune
	var threshold int

	if partB {
		neighbourFn = g.lineOfSightNeighbour
		threshold = 5
	} else {
		// part A
		neighbourFn = g.immediateNeighbour
		threshold = 4
	}

	if seat == 'L' {
		for _, dir := range dirs {
			if neighbourFn(c, dir) == '#' {
				return 0
			}
		}
		return 1
	} else if seat == '#' {
		adjOccupied := 0
		for _, dir := range dirs {
			if neighbourFn(c, dir) == '#' {
				adjOccupied++
			}
		}
		if adjOccupied >= threshold {
			return -1
		}
	}

	// floor - no change
	return 0
}

// returns '.' if go off the grid
func (g grid) get(c coord) rune {
	if c.row < 0 || c.row >= len(g) {
		return '.'
	}
	row := g[c.row]

	if c.col < 0 || c.col >= len(row) {
		return '.'
	}

	return row[c.col]
}

// returns first seat in specified direction, or floor if none found within the grid
func (g grid) first(c coord, dir coord) rune {
	for c.row >= 0 && c.col >= 0 && c.row < len(g) && c.col < len(g[0]) {
		c = add(c, dir)
		s := g.get(c)
		if s == '#' || s == 'L' {
			return s
		}
	}
	return '.'
}

func (g grid) toString() string {
	out := make([]rune, 0)

	for _, r := range [][]rune(g) {
		for _, c := range r {
			out = append(out, c)
		}
		out = append(out, '\n')
	}
	return string(out)
}

func (g grid) occupiedCount() int {
	count := 0
	for _, r := range g {
		for _, c := range r {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func (g grid) copy() grid {
	cpy := make(grid, len(g))
	for i, v := range g {
		rowcpy := make([]rune, len(v))
		copy(rowcpy, v)
		cpy[i] = rowcpy
	}
	return cpy
}

func (g grid) lineOfSightNeighbour(c, dir coord) rune {
	return g.first(c, dir)
}

func (g grid) immediateNeighbour(c, dir coord) rune {
	return g.get(add(c, dir))
}
