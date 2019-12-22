package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	inputFilePath = flag.String("input", "day12.txt", "input file path")
	steps         = flag.Int("steps", 1000, "number of timesteps to simulate")
	partB         = flag.Bool("partB", false, "enable part b")
)

type moon struct {
	position coord
	velocity coord
}

type coord [3]int // x, y, z

func main() {
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moons := readInput(file)

	process(moons, *steps, *partB)
}

func process(moons []moon, steps int, partB bool) {
	if partB {
		initial := make([]moon, 0, len(moons))
		for _, m := range moons {
			initial = append(initial, m)
		}

		intervals := make([]*int, 3)

		for t := 1; containsNil(intervals); t++ {
			doTimeStep(moons)

			for c := 0; c < 3; c++ {
				if intervals[c] == nil {
					if isSame(initial, moons, c) {
						t := t
						intervals[c] = &t
					}
				}
			}
		}

		ans := lcm(lcm(*intervals[0], *intervals[1]), *intervals[2])

		fmt.Printf("%d\n", ans)
	} else {
		// part A
		for t := 1; t <= steps; t++ {
			doTimeStep(moons)

			fmt.Printf("After %d steps:\n", t)
			for _, m := range moons {
				fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", m.position[0], m.position[1], m.position[2], m.velocity[0], m.velocity[1], m.velocity[2])
			}
		}

		totalEnergy := 0
		for _, m := range moons {
			e := m.kineticEnergy() * m.potentialEnergy()
			totalEnergy += e
		}

		fmt.Printf("%d\n", totalEnergy)
	}
}

func doTimeStep(moons []moon) {
	// update moon velocities
	for i := 0; i < len(moons)-1; i++ {
		for j := i; j < len(moons); j++ {
			// x
			updateVelocity(moons, i, j, 0)
			// y
			updateVelocity(moons, i, j, 1)
			// z
			updateVelocity(moons, i, j, 2)
		}
	}

	// apply velocity to update position
	for i, m := range moons {
		moons[i].position = m.position.add(m.velocity)
	}
}

func updateVelocity(moons []moon, i, j, c int) {
	if moons[i].position[c] < moons[j].position[c] {
		moons[i].velocity[c]++
		moons[j].velocity[c]--
	} else if moons[i].position[c] > moons[j].position[c] {
		moons[i].velocity[c]--
		moons[j].velocity[c]++
	}
}

func isSame(a []moon, b []moon, c int) bool {
	for i := range b {
		if !(b[i].position[c] == a[i].position[c] &&
			b[i].velocity[c] == a[i].velocity[c]) {
			return false
		}
	}
	return true
}

func readInput(r io.Reader) []moon {
	expr := regexp.MustCompile(`<x=(?P<x>[-\d]+),\sy=(?P<y>[-\d]+),\sz=(?P<z>[-\d]+)>`)
	scanner := bufio.NewScanner(r)
	moons := make([]moon, 0, 4)
	for scanner.Scan() {
		text := scanner.Text()
		match := expr.FindStringSubmatch(text)
		result := make(map[string]string)
		for i, name := range expr.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		x := toInt(result["x"])
		y := toInt(result["y"])
		z := toInt(result["z"])
		m := moon{position: coord{x, y, z}}
		moons = append(moons, m)
	}
	return moons
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func (c coord) add(other coord) coord {
	return coord{c[0] + other[0], c[1] + other[1], c[2] + other[2]}
}

func (m moon) potentialEnergy() int {
	p := m.position
	return abs(p[0]) + abs(p[1]) + abs(p[2])
}

func (m moon) kineticEnergy() int {
	v := m.velocity
	return abs(v[0]) + abs(v[1]) + abs(v[2])
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}

func containsNil(s []*int) bool {
	for _, v := range s {
		if v == nil {
			return true
		}
	}
	return false
}
