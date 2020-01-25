package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

var (
	inputFilePath = flag.String("input", "day20.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type coord struct {
	x int
	y int
}

type coordAndDepth struct {
	coord coord
	depth int
}

type portal string

func main() {
	flag.Parse()
	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parsedInput := readInput(file)

	ans := process(parsedInput)

	fmt.Println(ans)
}

func process(parsedInput [][]rune) int {
	portalToCoords := make(map[portal][]coord)
	coordToPortal := make(map[coord]portal)

	// Find portals
	for x := 2; x < len(parsedInput[0])-2; x++ {
		for y := 2; y < len(parsedInput)-2; y++ {
			c := parsedInput[y][x]
			if c != '.' {
				continue
			}

			coord := coord{x, y}

			if unicode.IsLetter(parsedInput[y-1][x]) {
				// portal is above
				p := portal(concat(parsedInput[y-2][x], parsedInput[y-1][x]))
				portalToCoords[p] = append(portalToCoords[p], coord)
				coordToPortal[coord] = p
			} else if unicode.IsLetter(parsedInput[y][x-1]) {
				p := portal(concat(parsedInput[y][x-2], parsedInput[y][x-1]))
				// portal is left
				portalToCoords[p] = append(portalToCoords[p], coord)
				coordToPortal[coord] = p
			} else if unicode.IsLetter(parsedInput[y+1][x]) {
				p := portal(concat(parsedInput[y+1][x], parsedInput[y+2][x]))
				// portal is below
				portalToCoords[p] = append(portalToCoords[p], coord)
				coordToPortal[coord] = p
			} else if unicode.IsLetter(parsedInput[y][x+1]) {
				p := portal(concat(parsedInput[y][x+1], parsedInput[y][x+2]))
				// portal is right
				portalToCoords[p] = append(portalToCoords[p], coord)
				coordToPortal[coord] = p
			}
		}
	}

	portalLinks := make(map[coord]coord, len(portalToCoords))

	// Connect portals
	for _, v := range portalToCoords {
		if len(v) > 1 { // only 1 portal for AA and ZZ
			portalLinks[v[0]] = v[1]
			portalLinks[v[1]] = v[0]
		}
	}

	// BFS
	start := portalToCoords["AA"][0]
	q := []coordAndDepth{coordAndDepth{coord: start}}

	visited := make(map[coord]bool)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if coordToPortal[curr.coord] == "ZZ" {
			// reached end
			return curr.depth
		}

		if v, ok := portalLinks[curr.coord]; ok {
			// portal link exists
			if _, ok := visited[portalLinks[curr.coord]]; !ok {
				// haven't gone through portal
				q = append(q, coordAndDepth{coord: v, depth: curr.depth + 1})
			}
		}

		for _, n := range neighbours(curr.coord) {
			if _, ok := visited[n]; !ok {
				// not visited neighbour
				if parsedInput[n.y][n.x] == '.' {
					q = append(q, coordAndDepth{coord: n, depth: curr.depth + 1})
				}
			}
		}
		visited[curr.coord] = true
	}
	return -1
}

// list of rows
func readInput(r io.Reader) [][]rune {
	lines := make([][]rune, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		cs := []rune(scanner.Text())
		lines = append(lines, cs)
	}
	return lines
}

func concat(r1, r2 rune) string {
	return string(r1) + string(r2)
}

func neighbours(c coord) []coord {
	return []coord{coord{c.x + 1, c.y}, coord{c.x - 1, c.y}, coord{c.x, c.y + 1}, coord{c.x, c.y - 1}}
}
