package main

import (
	"flag"
	"fmt"
)

var (
	endTurn = flag.Int("end", 2020, "turn to end game")
	numbers = []int{0, 8, 15, 2, 12, 1, 4}
)

func main() {
	flag.Parse()

	// spoken value to turns
	spoken := make(map[int][]int)

	for i, v := range numbers {
		spoken[v] = []int{i + 1}
	}

	lastSpoken := numbers[len(numbers)-1]

	for turn := len(numbers) + 1; turn <= *endTurn; turn++ {

		turns := spoken[lastSpoken]

		if len(turns) == 1 {
			spoken[0] = append(spoken[0], turn)
			lastSpoken = 0
		} else {
			mostRecent := turns[len(turns)-1]
			secondMostRecent := turns[len(turns)-2]
			diff := mostRecent - secondMostRecent

			if _, ok := spoken[diff]; !ok {
				spoken[diff] = make([]int, 0)
			}
			spoken[diff] = append(spoken[diff], turn)
			lastSpoken = diff
		}
	}

	fmt.Println(lastSpoken)
}
