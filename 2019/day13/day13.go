package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a-buck/adventofcode/2019/intcode"
)

var (
	inputFilePath = flag.String("input", "day13.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type coord struct {
	// distance from left
	x int
	// distance from the top
	y int
}

type tile int

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	program := intcode.ReadProgram(inputBytes)

	if *partB {
		program[0] = 2
	}

	joystick := make(chan int, 1)
	output := make(chan int, 1)

	go intcode.Run(program, joystick, output)

	paramNumber := 0
	var x int
	var y int
	var score int

	numberBlocks := 0

	ball := coord{}
	paddle := coord{}

	board := make(map[coord]tile)

	for i := range output {
		switch paramNumber {
		case 0:
			x = i
		case 1:
			y = i
		case 2:
			if x == -1 && y == 0 {
				score = i
			} else {
				c := coord{x, y}
				board[c] = tile(i)
				if i == 2 {
					// block tile
					numberBlocks++
				} else if i == 3 {
					// paddle
					paddle = c
				} else if i == 4 {
					// ball
					ball = c

					draw(board)

					fmt.Printf("ball: %v\n", ball)
					fmt.Printf("paddle: %v\n", paddle)

					if paddle.x < ball.x {
						// move right
						joystick <- 1
					} else if ball.x < paddle.x {
						// move left
						joystick <- -1
					} else {
						// stay
						joystick <- 0
					}
				}

			}

		}
		paramNumber++
		if paramNumber == 3 {
			paramNumber = 0
		}
	}
	if *partB {
		fmt.Println(score)
	} else {
		// part A
		fmt.Println(numberBlocks)
	}
}

func draw(board map[coord]tile) {
	maxX := 0
	maxY := 0

	for k := range board {
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			v := board[coord{x, y}]
			if v == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", v)
			}
		}
		fmt.Printf("\n")
	}
}
