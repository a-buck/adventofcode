package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	inputFilePath = flag.String("input", "day03.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type coord struct {
	x int
	y int
}

type grid [][]rune

func main() {
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := run(file, *partB)

	fmt.Println(ans)
}

func run(r io.Reader, partB bool) int {
	scanner := bufio.NewScanner(r)
	grid := grid(make([][]rune, 0))

	for scanner.Scan() {
		row := []rune(scanner.Text())
		grid = append(grid, row)
	}

	a := calcTreesHit(grid, coord{x: 3, y: 1})
	if partB {
		b := calcTreesHit(grid, coord{x: 1, y: 1})
		c := calcTreesHit(grid, coord{x: 5, y: 1})
		d := calcTreesHit(grid, coord{x: 7, y: 1})
		e := calcTreesHit(grid, coord{x: 1, y: 2})
		return a * b * c * d * e
	} else {
		// part A
		return a
	}
}

func calcTreesHit(g grid, c coord) int {
	curr := coord{}
	treesHit := 0

	for curr.y < len(g) {
		r := g.get(curr)
		if r == '#' {
			treesHit++
		}
		curr = curr.add(c)
	}
	return treesHit
}

func (g grid) get(c coord) rune {
	return g[c.y][c.x%len(g[0])]
}

func (c coord) add(c2 coord) coord {
	return coord{x: c.x + c2.x, y: c.y + c2.y}
}
