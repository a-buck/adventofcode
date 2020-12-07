package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type edge struct {
	destination string
	weight      int
}

var (
	inputFilePath = flag.String("input", "day07.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	ans := run(string(content), *partB)

	fmt.Println(ans)
}

func run(content string, partB bool) int {

	lines := strings.Split(content, "\n")

	adjlist := make(map[string][]edge, len(lines))

	for _, line := range lines {
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")

		lineParts := strings.Split(line, " contain ")
		parentBag := strings.TrimSpace(lineParts[0])

		childBags := regexp.MustCompile("[.,]").Split(lineParts[1], -1)

		childBagsTrimmed := make([]string, 0, len(childBags))
		for _, c := range childBags {
			trimmed := strings.TrimSpace(c)
			if trimmed != "" {
				childBagsTrimmed = append(childBagsTrimmed, trimmed)
			}
		}

		adjlist[parentBag] = make([]edge, 0)

		for _, b := range childBagsTrimmed {
			if b == "no other" {
				continue
			}
			amount, err := strconv.Atoi(b[0:1])
			if err != nil {
				log.Fatal(err)
			}
			childName := b[2:]
			e := edge{childName, amount}
			adjlist[parentBag] = append(adjlist[parentBag], e)
		}
	}

	if partB {
		return totalBags("shiny gold", adjlist)
	}

	// part A
	return goldReachableCount(adjlist)

}

func goldReachableCount(adjlist map[string][]edge) int {
	count := 0

	for source := range adjlist {
		if source == "shiny gold" {
			continue
		}

		if canReachGold(source, adjlist) {
			count++
		}
	}
	return count
}

func canReachGold(source string, adjlist map[string][]edge) bool {
	for _, n := range adjlist[source] {
		if n.destination == "shiny gold" || canReachGold(n.destination, adjlist) {
			return true
		}
	}
	return false
}

func totalBags(source string, adjlist map[string][]edge) int {
	sum := 0
	for _, n := range adjlist[source] {
		sum += n.weight*totalBags(n.destination, adjlist) + n.weight
	}
	return sum
}
