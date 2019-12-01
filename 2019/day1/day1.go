package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"fmt"
)

func calcFuel(mass int) int {
	return mass/3 - 2
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

	scanner := bufio.NewScanner(file)

	fuel := 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		var f int
		if *partB {
			f = mass
			for {
				f = calcFuel(f)
				if f <= 0 {
					break
				}
				fuel += f
			}


		} else {
			f = calcFuel(mass)
			fuel += f
		}

	}

	fmt.Println(fuel)

}
