package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var inputFilePath = flag.String("input", "day6.txt", "input file path")
	var partB = flag.Bool("partB", false, "enable part b")
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans, err := run(file, *partB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}

func run(inputReader io.Reader, partB bool) (int, error) {
	scanner := bufio.NewScanner(inputReader)

	adjlist := make(map[string][]string)

	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, ")")

		if len(tokens) != 2 {
			return 0, fmt.Errorf("got %d tokens, expected 2", len(tokens))
		}

		key := tokens[0]
		val := tokens[1]

		adjlist[key] = append(adjlist[key], val)
		adjlist[val] = append(adjlist[val], key)
	}

	var startNode string

	if partB {
		startNode = "YOU"
	} else {
		// part A
		startNode = "COM"
	}

	depths := traverse(startNode, adjlist)

	if partB {
		return depths["SAN"] - 2, nil
	}

	// part A
	depthTotal := 0
	for _, d := range depths {
		depthTotal += d
	}
	return depthTotal, nil
}

func traverse(root string, adjlist map[string][]string) map[string]int {
	depths := make(map[string]int, len(adjlist))
	nodeQueue := []string{root}
	depthQueue := []int{0}

	for len(nodeQueue) > 0 {
		curr := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		depth := depthQueue[0]
		depths[curr] = depth
		depthQueue = depthQueue[1:]

		for _, neighbour := range adjlist[curr] {
			_, visited := depths[neighbour]

			if !visited {
				nodeQueue = append(nodeQueue, neighbour)
				depthQueue = append(depthQueue, depth+1)
			}
		}
	}
	return depths
}
