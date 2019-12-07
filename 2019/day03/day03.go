package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func main() {
	var inputFilePath = flag.String("input", "day1.txt", "input file path")
	var partB = flag.Bool("partB", false, "enable part b")

	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := run(file, *partB)

	fmt.Println(result)
}

func run(reader io.Reader, partB bool) int {
	scanner := bufio.NewScanner(reader)

	coordToMinDistList := make([]map[coord]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coordToDist := make(map[coord]int)
		instructions := strings.Split(line, ",")

		currentX := 0
		currentY := 0

		dist := 0

		for _, inst := range instructions {
			direction := string(inst[0])
			amount, err := strconv.Atoi(string(inst[1:]))
			if err != nil {
				log.Fatal(err)
			}

			dx := 0
			dy := 0

			switch direction {
			case "R":
				dx = 1
			case "L":
				dx = -1
			case "U":
				dy = 1
			case "D":
				dy = -1
			}

			for i := 0; i < amount; i++ {
				dist++
				currentX += dx
				currentY += dy

				c := coord{currentX, currentY}
				_, ok := coordToDist[c]
				if !ok {
					coordToDist[c] = dist
				}
			}
		} // end of instruction

		coordToMinDistList = append(coordToMinDistList, coordToDist)
	} // end of line

	smallest := math.MaxInt64

	for c1, dist := range coordToMinDistList[0] {
		for c2, dist2 := range coordToMinDistList[1] {
			if c1 == c2 {

				if partB {
					totaldist := dist + dist2
					if totaldist < smallest {
						smallest = totaldist
					}
				} else {
					manhattenDistance := abs(c1.x) + abs(c1.y)
					if manhattenDistance < smallest {
						smallest = manhattenDistance
					}
				}
			}

		}
	}
	return smallest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
