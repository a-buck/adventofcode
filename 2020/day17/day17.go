package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day17.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type coord interface {
	neighbours() []coord
}

type coord4d struct {
	x, y, z, w int
}

type coord3d struct {
	x, y, z int
}

type coordState map[coord]rune

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	state := make(coordState)

	for y, line := range strings.Split(string(content), "\n") {
		for x, v := range line {
			if v == '#' {
				var c coord
				if *partB {
					c = coord4d{x: x, y: y}
				} else {
					c = coord3d{x: x, y: y}
				}
				state[c] = '#'
			}
		}
	}

	for i := 0; i < 6; i++ {
		seen := make(map[coord]bool)

		coordToActiveCount := make(map[coord]int)

		q := make([]coord, 0)

		for k, v := range state {
			if v == '#' {
				q = append(q, k)
			}
		}

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if _, ok := seen[curr]; ok {
				continue
			}

			for _, n := range curr.neighbours() {
				if _, ok := seen[n]; !ok {
					if state.get(n) == '#' {
						q = append(q, n)
					}

				}

				if state.get(curr) == '#' {
					coordToActiveCount[n]++

				}
			}

			seen[curr] = true
		}

		for k, v := range coordToActiveCount {
			if !(state.get(k) == '#' && (v == 2 || v == 3)) {
				delete(state, k)
			}

			if state.get(k) == '.' && v == 3 {
				state[k] = '#'
			}
		}

		for k, v := range state {
			if v == '.' {
				continue
			}
			if _, ok := coordToActiveCount[k]; !ok {
				delete(state, k)
			}
		}

	}

	activeCount := 0
	for _, v := range state {
		if v == '#' {
			activeCount++
		}

	}
	fmt.Println(activeCount)

}

func (c coord4d) neighbours() []coord {
	neighbours := make([]coord, 0, 80)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				for w := c.w - 1; w <= c.w+1; w++ {
					n := coord4d{x, y, z, w}
					if c != n {
						neighbours = append(neighbours, n)
					}
				}
			}
		}

	}
	return neighbours
}

func (c coord3d) neighbours() []coord {
	neighbours := make([]coord, 0, 26)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				n := coord3d{x, y, z}
				if c != n {
					neighbours = append(neighbours, n)
				}
			}
		}

	}
	return neighbours
}

func (g coordState) get(c coord) rune {
	if v, ok := g[c]; ok {
		return v
	}
	return '.'
}
