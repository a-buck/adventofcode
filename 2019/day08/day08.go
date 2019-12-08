package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var (
	inputFilePath = flag.String("input", "day08.txt", "input file path")
	width         = flag.Int("width", 25, "image width")
	height        = flag.Int("height", 6, "image height")
)

func main() {
	flag.Parse()

	inputBytes, err := ioutil.ReadFile(*inputFilePath)

	if err != nil {
		log.Fatal(err)
	}

	pixels := strings.Split(string(inputBytes), "")

	process(pixels, *width, *height)
}

func process(pixels []string, width int, height int) {
	minZeros := math.MaxInt32
	partAResult := 0
	pIdx := 0

	img := make([]string, width*height)

	for layer := 0; pIdx < len(pixels); layer++ {
		layerCounts := make(map[string]int)

		for row := 0; row < height; row++ {

			for col := 0; col < width; col, pIdx = col+1, pIdx+1 {
				pixel := pixels[pIdx]

				layerCounts[pixel]++

				imgIdx := index(row, col, width)
				curr := img[imgIdx]

				if curr == "" || curr == "2" {
					// Overwrite if existing pixel is empty or transparent (2)
					img[imgIdx] = pixel
				}
			}
		}

		zeroCount := layerCounts["0"]
		if zeroCount < minZeros {
			minZeros = zeroCount
			partAResult = layerCounts["1"] * layerCounts["2"]
		}

	}

	fmt.Println(partAResult)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {

			imgIdx := img[index(row, col, width)]

			if imgIdx == "1" {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n")
	}
}

func index(row int, col int, width int) int {
	return row*width + col
}
