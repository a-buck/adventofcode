package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"unicode"

	"github.com/a-buck/adventofcode/2019/priorityqueue"
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

type coordAndLevel struct {
	coord coord
	level int
}

type edge struct {
	n          coord
	dist       int
	levelDelta int
}

type portal struct {
	name    string
	isOuter bool // if not outer, it is inner
}

type portalAndCoord struct {
	portal portal
	coord  coord
}

func main() {
	flag.Parse()
	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parsedInput := readInput(file)

	ans := process(parsedInput, *partB)

	fmt.Println(ans)
}

func process(parsedInput [][]rune, partB bool) int {
	portalAndCoords := getPortals(parsedInput)

	portalToCoords := createPortalToCoords(portalAndCoords)
	coordToPortal := createCoordToPortal(portalAndCoords)
	portalLinks := getPortalLinks(portalToCoords)

	adjlist := createAdjlist(coordToPortal, parsedInput)

	dists := dijkstra(portalToCoords, coordToPortal, adjlist, portalLinks, partB)
	zzCoord := portalToCoords["ZZ"][0]
	if v, ok := dists[coordAndLevel{zzCoord, 0}]; ok {
		return v
	}
	return -1
}

func dijkstra(portalToCoords map[string][]coord, coordToPortal map[coord]portal, adjlist map[coord][]edge, portalLinks map[coord]coord, partB bool) map[coordAndLevel]int {
	dists := make(map[coordAndLevel]int)
	visited := make(map[coordAndLevel]bool)

	q := make(priorityqueue.MinPriorityQueue, 1)
	start := portalToCoords["AA"][0]
	q[0] = &priorityqueue.Item{Value: coordAndLevel{start, 0}}

	heap.Init(&q)

	for len(q) > 0 {
		item := heap.Pop(&q).(*priorityqueue.Item)
		curr := item.Value.(coordAndLevel)

		// add linked portal as neighbour
		neighbours := adjlist[curr.coord]
		if v, ok := portalLinks[curr.coord]; ok {
			currPortal := coordToPortal[curr.coord]
			levelDelta := toLevelDelta(currPortal.isOuter, partB)
			neighbours = append(neighbours, edge{v, 1, levelDelta})
		}

		if coordToPortal[curr.coord].name == "ZZ" && curr.level == 0 {
			return dists
		}

		for _, neighbourEdge := range neighbours {
			neighbourCoordAndLevel := coordAndLevel{neighbourEdge.n, curr.level + neighbourEdge.levelDelta}

			if neighbourCoordAndLevel.level < 0 {
				continue
			}

			if _, ok := visited[neighbourCoordAndLevel]; ok {
				// already visited neighbour
				continue
			}

			// init dist for neighbour to infinity if not initialised.
			if _, ok := dists[neighbourCoordAndLevel]; !ok {
				dists[neighbourCoordAndLevel] = math.MaxInt32
			}

			if dists[curr]+neighbourEdge.dist < dists[neighbourCoordAndLevel] {
				dists[neighbourCoordAndLevel] = dists[curr] + neighbourEdge.dist
			}

			q.Push(&priorityqueue.Item{Value: neighbourCoordAndLevel, Priority: dists[neighbourCoordAndLevel]})
		}

		visited[curr] = true
	}
	return dists
}

func toLevelDelta(isOuter bool, partB bool) int {
	if !partB {
		return 0
	}
	if isOuter {
		return -1
	}
	return 1
}

// getPortal if exists, else nil
func getPortal(parsedInput [][]rune, x int, y int) *portal {
	var p *portal
	if unicode.IsLetter(parsedInput[y-1][x]) {
		// portal is above
		p = &portal{name: concat(parsedInput[y-2][x], parsedInput[y-1][x]), isOuter: y == 2}
	} else if unicode.IsLetter(parsedInput[y][x-1]) {
		// portal is left
		p = &portal{name: concat(parsedInput[y][x-2], parsedInput[y][x-1]), isOuter: x == 2}
	} else if unicode.IsLetter(parsedInput[y+1][x]) {
		// portal is below
		p = &portal{name: concat(parsedInput[y+1][x], parsedInput[y+2][x]), isOuter: y == len(parsedInput)-3}
	} else if unicode.IsLetter(parsedInput[y][x+1]) {
		// portal is right
		p = &portal{name: concat(parsedInput[y][x+1], parsedInput[y][x+2]), isOuter: x == len(parsedInput[y])-3}
	}
	return p
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

func getPortals(parsedInput [][]rune) []portalAndCoord {
	portalAndCoords := make([]portalAndCoord, 0)

	for x := 2; x < len(parsedInput[0])-2; x++ {
		for y := 2; y < len(parsedInput)-2; y++ {
			c := parsedInput[y][x]
			if c != '.' {
				continue
			}

			coord := coord{x, y}

			p := getPortal(parsedInput, x, y)

			if p != nil {
				portalAndCoords = append(portalAndCoords, portalAndCoord{*p, coord})
			}
		}
	}
	return portalAndCoords
}

func getPortalLinks(portalToCoords map[string][]coord) map[coord]coord {
	portalLinks := make(map[coord]coord, len(portalToCoords))

	// Connect portals
	for p, v := range portalToCoords {

		if len(v) > 1 { // only 1 portal for AA and ZZ
			portalLinks[v[0]] = v[1]
			portalLinks[v[1]] = v[0]
		} else if len(v) <= 1 {
			if p != "AA" && p != "ZZ" {
				log.Fatalf("Tried to connect portal %s. That should never happen", p)
			}
		}
	}
	return portalLinks
}

func createPortalToCoords(portalAndCoords []portalAndCoord) map[string][]coord {
	portalToCoords := make(map[string][]coord)
	for _, p := range portalAndCoords {
		portalToCoords[p.portal.name] = append(portalToCoords[p.portal.name], p.coord)
	}
	return portalToCoords
}

func createCoordToPortal(portalAndCoords []portalAndCoord) map[coord]portal {
	coordToPortal := make(map[coord]portal)
	for _, p := range portalAndCoords {
		coordToPortal[p.coord] = p.portal
	}
	return coordToPortal
}

func createAdjlist(coordToPortal map[coord]portal, parsedInput [][]rune) map[coord][]edge {
	adjlist := make(map[coord][]edge)

	for c := range coordToPortal {
		visited := make(map[coord]bool, 0)

		q := []coordAndDepth{coordAndDepth{coord: c}}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if _, ok := coordToPortal[curr.coord]; ok {
				// reached a portal
				if curr.coord != c {
					adjlist[c] = append(adjlist[c], edge{curr.coord, curr.depth, 0})
				}
			}

			for _, n := range neighbours(curr.coord) {
				if parsedInput[n.y][n.x] == '.' {
					if _, ok := visited[coord{n.x, n.y}]; !ok {
						q = append(q, coordAndDepth{n, curr.depth + 1})
					}
				}
			}
			visited[curr.coord] = true
		}
	}
	return adjlist
}
