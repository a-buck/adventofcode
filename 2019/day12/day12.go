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

type coord struct {
	x int
	y int
	z int
}

func main() {
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moons := readInput(file)

	// fmt.Printf("%+v\n", moons)

	process(moons, *steps, *partB)

}

type fourmoons struct {
	m1 moon
	m2 moon
	m3 moon
	m4 moon
}

func process(moons []moon, steps int, partB bool) {
	if partB {
		seen := make(map[fourmoons]bool)

		t := 0
		for ; ; t++ {
			doTimeStep(moons)
			fm := fourmoons{moons[0], moons[1], moons[2], moons[3]}
			_, ok := seen[fm]
			if ok {
				break
			} else {
				seen[fm] = true
			}
		}
		fmt.Printf("Steps: %d\n", t)
	} else {
		// part A
		for t := 1; t <= steps; t++ {
			doTimeStep(moons)

			fmt.Printf("After %d steps:\n", t)
			for _, m := range moons {
				fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", m.position.x, m.position.y, m.position.z, m.velocity.x, m.velocity.y, m.velocity.z)
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
	for i := 0; i < len(moons)-1; i++ {
		for j := i; j < len(moons); j++ {
			m1 := moons[i]
			m2 := moons[j]

			// x
			if m1.position.x < m2.position.x {
				moons[i].velocity.x++
				moons[j].velocity.x--
			} else if m1.position.x > m2.position.x {
				moons[i].velocity.x--
				moons[j].velocity.x++
			}

			// y
			if m1.position.y < m2.position.y {
				moons[i].velocity.y++
				moons[j].velocity.y--
			} else if m1.position.y > m2.position.y {
				moons[i].velocity.y--
				moons[j].velocity.y++
			}

			// z
			if m1.position.z < m2.position.z {
				moons[i].velocity.z++
				moons[j].velocity.z--
			} else if m1.position.z > m2.position.z {
				moons[i].velocity.z--
				moons[j].velocity.z++
			}

		}
	}

	// 2) apply velocity to update position
	for i, m := range moons {
		moons[i].position = m.position.add(m.velocity)
	}
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
		m := moon{position: coord{x: x, y: y, z: z}}
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
	return coord{x: c.x + other.x, y: c.y + other.y, z: c.z + other.z}
}

func (m moon) potentialEnergy() int {
	p := m.position
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func (m moon) kineticEnergy() int {
	v := m.velocity
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
