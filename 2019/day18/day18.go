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
	inputFilePath = flag.String("input", "day18.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type coord struct {
	x int
	y int
}

type adjlist map[coord][]coord

type edge struct {
	dist      int
	neighbour coord
}

type coordAndKeys struct {
	c        coord
	keysMask int
}

type coordAndDepth struct {
	c coord
	d int
}

func main() {
	flag.Parse()
	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	coords := toCoords(file)
	adjlist := toAdjList(coords)
	compressedAdjlist := toCompressedAdjlist(adjlist, coords)

	ans := process(compressedAdjlist, coords, *partB)

	fmt.Println(ans)
}

func toCoords(r io.Reader) map[coord]rune {
	scanner := bufio.NewScanner(r)

	coords := make(map[coord]rune)

	for row := 0; scanner.Scan(); row++ {
		cs := []rune(scanner.Text())
		for col, c := range cs {
			coords[coord{col, row}] = c
		}
	}
	return coords
}

func toAdjList(coords map[coord]rune) adjlist {

	adjlist := make(adjlist, len(coords))

	deltas := []coord{{y: -1}, {x: 1}, {y: 1}, {x: -1}} // up, right, down, left

	for k, v := range coords {
		if v == '#' {
			continue
		}
		for _, d := range deltas {
			n := k.add(d)
			if coords[n] != '#' {
				adjlist[k] = append(adjlist[k], n)
			}
		}
	}
	return adjlist
}

func toCompressedAdjlist(adjlist adjlist, coords map[coord]rune) map[coord][]edge {
	compressedAdjList := make(map[coord][]edge)

	for root, r := range coords {
		if r == '#' || r == '.' {
			continue
		}

		// do BFS to get depth of nearest neighbours
		visited := make(map[coord]bool)
		q := make([]coordAndDepth, 0)
		q = append(q, coordAndDepth{c: root})

		for len(q) > 0 {

			curr := q[0]
			q = q[1:]

			if coords[curr.c] != '#' && coords[curr.c] != '.' && curr.c != root { // not empty space or wall
				compressedAdjList[root] = append(compressedAdjList[root], edge{curr.d, curr.c})
			} else {
				for _, n := range adjlist[curr.c] {
					if _, ok := visited[n]; !ok {
						q = append(q, coordAndDepth{n, curr.d + 1})
					}
				}
			}
			visited[curr.c] = true
		}
	}

	return compressedAdjList

}

func process(compressedAdjList map[coord][]edge, coords map[coord]rune, partB bool) int {

	var roots = make([]coord, 0)

	keysBitMask := 0

	for k, v := range coords {
		if v == '@' {
			roots = append(roots, k)
		} else if isLowercaseLetter(v) {
			if mask, err := set(v, keysBitMask); err != nil {
				log.Fatal(err)
			} else {
				keysBitMask = mask
			}
		}
	}

	if partB {
		total := 0
		for _, r := range roots {
			visited := make(map[coord]bool)
			// generate  bitmask for keys accessible from this root.
			localBitMask := 0
			q := make([]coord, 0)
			q = append(q, r)
			for len(q) > 0 {
				curr := q[0]
				q = q[1:]

				if isLowercaseLetter(coords[curr]) { // is key
					var err error
					localBitMask, err = set(coords[curr], localBitMask)
					if err != nil {
						log.Fatal(err)
					}
				}

				for _, n := range compressedAdjList[curr] {
					if _, ok := visited[n.neighbour]; !ok {
						q = append(q, n.neighbour)
					}
				}
				visited[curr] = true
			}

			dist := distToGetAllKeys(compressedAdjList, coords, localBitMask, r)
			total += dist
		}
		return total

	} else {
		// part A
		dist := distToGetAllKeys(compressedAdjList, coords, keysBitMask, roots[0])
		return dist
	}
}

// dijkstra to get dist for all keys
func distToGetAllKeys(compressedAdjList map[coord][]edge, coords map[coord]rune, keysBitMask int, root coord) int {
	dists := make(map[coordAndKeys]int, len(compressedAdjList))
	visited := make(map[coordAndKeys]bool, len(compressedAdjList))

	remaining := make(priorityqueue.MinPriorityQueue, 1)

	remaining[0] = &priorityqueue.Item{Value: coordAndKeys{root, 0}}

	heap.Init(&remaining)

	for len(remaining) > 0 {
		item := heap.Pop(&remaining).(*priorityqueue.Item)
		curr := item.Value.(coordAndKeys)

		initialCurr := curr

		val := coords[curr.c]

		if isLowercaseLetter(val) { // is key
			mask, err := set(val, curr.keysMask)
			curr.keysMask = mask
			dists[curr] = dists[initialCurr]
			if err != nil {
				log.Fatal(err)
			}
		} else if isUppercaseLetter(val) {
			// is door
			if keyAccessible, _ := isSet(unicode.ToLower(val), keysBitMask); keyAccessible {
				// key is accessible
				if isSet, _ := isSet(unicode.ToLower(val), curr.keysMask); !isSet {
					// don't have key
					continue
				}
			}
		}

		if curr.keysMask == keysBitMask {
			return dists[curr]
		}

		for _, neighbourEdge := range compressedAdjList[curr.c] {
			neighbourCoordAndKey := coordAndKeys{neighbourEdge.neighbour, curr.keysMask}

			// skip if already visited neighbour with currently held keys
			if _, ok := visited[neighbourCoordAndKey]; ok {
				continue
			}

			// init dist for neighbour to infinity if not initialised.
			if _, ok := dists[neighbourCoordAndKey]; !ok {
				dists[neighbourCoordAndKey] = math.MaxInt32
			}

			// see if current route is shorter.
			if dists[curr]+neighbourEdge.dist < dists[neighbourCoordAndKey] {
				dists[neighbourCoordAndKey] = dists[curr] + neighbourEdge.dist

				// add neighbour to queue.
				remaining.Push(&priorityqueue.Item{Value: coordAndKeys{neighbourEdge.neighbour, curr.keysMask}, Priority: dists[neighbourCoordAndKey]})
			}

		}
		visited[curr] = true
	}
	return -1
}

func (c coord) add(other coord) coord {
	return coord{c.x + other.x, c.y + other.y}
}

func isLowercaseLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUppercaseLetter(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func set(r rune, mask int) (int, error) {
	if r < 'a' || r > 'z' {
		return 0, fmt.Errorf("Rune %c is out of range", r)
	}
	return mask | 1<<int(r-'a'), nil
}

func isSet(r rune, mask int) (bool, error) {
	if r < 'a' || r > 'z' {
		return false, fmt.Errorf("Rune %c is out of range", r)
	}
	bit := 1 << int(r-'a')
	isSet := bit&mask == bit
	return isSet, nil
}
