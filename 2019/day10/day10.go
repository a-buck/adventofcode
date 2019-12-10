package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var inputFilePath = flag.String("input", "day10.txt", "input file path")

type radians float64

type coord struct {
	// distance from left edge
	x int
	// distance from the top edge
	y int
}

func main() {
	flag.Parse()
	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	asteroids := parseInput(file)

	process(asteroids)
}

func process(asteroids []coord) {
	partAAns, maxCountAst := partA(asteroids)
	partBAns, _ := partB(asteroids, maxCountAst)

	fmt.Printf("part A ans: %d\n", partAAns)
	fmt.Printf("part B ans: %d\n", partBAns)
}

func parseInput(r io.Reader) []coord {
	asteroids := make([]coord, 0)
	scanner := bufio.NewScanner(r)

	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()
		tokens := strings.Split(text, "")
		for x, v := range tokens {
			if v == "#" {
				ast := coord{x: x, y: y}
				asteroids = append(asteroids, ast)
			}
		}
	}
	return asteroids
}

func partA(asteroids []coord) (maxVisibleAsteroids int, bestAsteroid coord) {

	for _, root := range asteroids {
		seenAngles := make(map[radians]bool, 0)

		count := 0

		for _, dest := range asteroids {
			if root == dest {
				continue
			}

			lineAngle := calcAngle(root, dest)

			_, ok := seenAngles[lineAngle]
			if !ok {
				seenAngles[lineAngle] = true
				count++
			}
		}

		if count > maxVisibleAsteroids {
			maxVisibleAsteroids = count
			bestAsteroid = root
		}
	}

	return maxVisibleAsteroids, bestAsteroid
}

func partB(asteroids []coord, root coord) (int, coord) {
	angleNeighbours := make(map[radians][]coord)

	// Group all coordinates by angle from root
	for _, ast := range asteroids {
		if ast == root {
			continue
		}
		angle := calcAngle(root, ast)
		angleNeighbours[angle] = append(angleNeighbours[angle], ast)
	}

	ascAngles := getKeysAsc(angleNeighbours)
	// sort asteroids by distance asc for each angle
	for angle, neighbourCoords := range angleNeighbours {
		sortedCoordsPerAngle := sortCoordsByDist(root, neighbourCoords)
		angleNeighbours[angle] = sortedCoordsPerAngle
	}

	var vaporisedAst200 coord
	vaporisedCount := 0
	currIndex := 0

	for vaporisedCount < len(asteroids)-1 {
		anglesWithNeighbours := ascAngles[currIndex]
		neighbours := angleNeighbours[anglesWithNeighbours]
		if len(neighbours) > 0 {
			vaporised := neighbours[0]
			angleNeighbours[anglesWithNeighbours] = neighbours[1:]
			vaporisedCount++
			if vaporisedCount == 200 {
				vaporisedAst200 = vaporised
			}
		}
		if currIndex < len(ascAngles)-1 {
			currIndex++
		} else {
			currIndex = 0
		}
	}

	partBAns := vaporisedAst200.x*100 + vaporisedAst200.y
	return partBAns, vaporisedAst200
}

func getKeysAsc(m map[radians][]coord) []radians {
	rs := make([]radians, 0)
	for k := range m {
		rs = append(rs, k)
	}

	rSorted := sortRadians(rs)
	return rSorted
}

func sortCoordsByDist(root coord, coords []coord) []coord {
	sort.Slice(coords, func(i, j int) bool {
		c1 := coords[i]
		c2 := coords[j]
		return dist(root, c1) < dist(root, c2)
	})

	return coords
}

func sortRadians(rs []radians) []radians {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	return rs

}

func calcAngle(c1, c2 coord) radians {
	// cosine rule
	// c2 = a2 + b2 - 2ab cos(C)

	refCoord := coord{c1.x, c1.y - 1} // y coord doesn't matter - just needs to be < c1.y

	c := dist(refCoord, c2)
	b := float64(c1.y - refCoord.y)
	a := dist(c1, c2)

	C := math.Acos((sq(a) + sq(b) - sq(c)) / (2 * a * b))

	if c2.x < c1.x {
		return radians(round(2*math.Pi - C))
	}
	return radians(round(C))
}

func dist(c1, c2 coord) float64 {
	return math.Sqrt(sq(float64(c2.x-c1.x)) + sq(float64(c2.y-c1.y)))
}

func round(f float64) float64 {
	return math.Round(f*10000) / 10000
}

func sq(f float64) float64 {
	return math.Pow(f, 2)
}
