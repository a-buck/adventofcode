package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type chemical struct {
	name   string
	amount int
}

type equation struct {
	inputs []chemical
	output chemical
}

type edge struct {
	source string
	dest   string
}

const oneTrillion = 1000000000000

var (
	inputFilePath = flag.String("input", "day14.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {

	flag.Parse()
	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	eqns := parseInput(file)
	eqns = reverseTopologicalSortEquations(eqns)

	if *partB {
		// binary search to find max amount of fuel
		// that can be produced with 1 trillion ore
		r := oneTrillion
		l := 0

		maxFuel := 0

		for l < r {
			mid := (l + r) / 2
			ore := process(eqns, mid)
			if ore <= oneTrillion {
				maxFuel = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		fmt.Println(maxFuel)

	} else {
		// part A
		ore := process(eqns, 1)
		fmt.Println(ore)
	}
}

// calculates how much ore is needed to create given amount of fuel
func process(equations []equation, fuel int) int {
	chemicalAmounts := make(map[string]int) // chemical to current amount held
	chemicalAmounts["FUEL"] = -1 * fuel

	for _, eqn := range equations {
		currentAmount := chemicalAmounts[eqn.output.name]
		numberReactions := 0

		// if there is a 'debt', need to do some reactions to get amount >= 0
		if currentAmount < 0 {
			numberReactions = -currentAmount / eqn.output.amount
			// if there is a remainder, do one more reaction to cover this
			if currentAmount%eqn.output.amount != 0 {
				numberReactions++
			}
		}

		// do reaction - subtract amount held for inputs
		// and add amount held for outputs
		for _, input := range eqn.inputs {
			chemicalAmounts[input.name] -= input.amount * numberReactions
		}
		chemicalAmounts[eqn.output.name] += (eqn.output.amount * numberReactions)
	}

	return chemicalAmounts["ORE"] * -1
}

func parseInput(r io.Reader) []equation {
	scanner := bufio.NewScanner(r)

	equations := make([]equation, 0)

	for scanner.Scan() {

		eqn := equation{}
		text := scanner.Text()
		eqnSplit := strings.Split(text, "=>")

		lhs := strings.Trim(eqnSplit[0], " ")
		lhsParts := strings.Split(lhs, ",")

		rhs := strings.Trim(eqnSplit[1], " ")
		rhsParts := strings.Split(rhs, " ")
		rhsChemical := rhsParts[1]
		rhsAmount := toInt(rhsParts[0])

		eqn.output.name = rhsChemical
		eqn.output.amount = rhsAmount

		for _, lhsPart := range lhsParts {
			lhsPart = strings.Trim(lhsPart, " ")
			lhsPartSplit := strings.Split(lhsPart, " ")
			amount := toInt(lhsPartSplit[0])
			chem := lhsPartSplit[1]

			eqn.inputs = append(eqn.inputs, chemical{chem, amount})
		}
		equations = append(equations, eqn)
	}
	return equations
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func reverseTopologicalSortEquations(eqns []equation) []equation {
	edges := toEdges(eqns)
	graph := toAdjList(edges)
	sortedNodes := reverseTopologicalSort(graph, "ORE")
	sortedEquations := toEquations(sortedNodes, eqns)
	return sortedEquations
}

// map chemicals to their corresponding equations
func toEquations(nodes []string, eqns []equation) []equation {
	eqnsByOutput := make(map[string]equation)
	for _, x := range eqns {
		eqnsByOutput[x.output.name] = x
	}

	sortedEquations := make([]equation, 0, len(eqns))
	for _, o := range nodes {
		eqn, ok := eqnsByOutput[o]
		if ok {
			sortedEquations = append(sortedEquations, eqn)
		}
	}
	return sortedEquations
}

func toEdges(eqns []equation) []edge {
	edges := make([]edge, 0, len(eqns))
	for _, eqn := range eqns {
		for _, i := range eqn.inputs {
			e := edge{i.name, eqn.output.name}
			edges = append(edges, e)
		}
	}
	return edges
}

func toAdjList(edges []edge) map[string][]string {
	adjlist := make(map[string][]string)

	for _, e := range edges {
		adjlist[e.source] = append(adjlist[e.source], e.dest)
	}
	return adjlist
}

func reverseTopologicalSort(adjlist map[string][]string, root string) []string {
	sorted := make([]string, 0, len(adjlist))
	visited := make(map[string]bool)
	reverseTopoSortVisit(adjlist, root, visited, &sorted)
	return sorted
}

func reverseTopoSortVisit(graph map[string][]string, node string, visited map[string]bool, sorted *[]string) {
	for _, neighbour := range graph[node] {
		if _, ok := visited[neighbour]; !ok {
			reverseTopoSortVisit(graph, neighbour, visited, sorted)
		}
	}
	visited[node] = true
	*sorted = append(*sorted, node)
}
