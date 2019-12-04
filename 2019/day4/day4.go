package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var minVal = flag.Int("min", 265275, "min possible value")
	var maxVal = flag.Int("max", 781584, "max possible value")
	var partB = flag.Bool("partB", false, "enable part b")

	flag.Parse()

	count := 0

	for i := *minVal; i <= *maxVal; i++ {
		ok := evaluateCandidate(i, *partB)
		if ok {
			count++
		}
	}
	fmt.Println(count)

}

func evaluateCandidate(candidate int, partB bool) bool {
	iStr := strconv.Itoa(candidate)

	adjEqCount := 0
	doubleFound := false

	prev := iStr[0]

	// check if number of sequential equal characters is valid
	// varies between part A/B
	accCheck := func(acc int) bool {
		if partB {
			return acc == 1
		}
		return acc >= 1
	}

	for j := 1; j < len(iStr); j++ {
		if iStr[j] == prev {
			// character same as previous
			adjEqCount++
		} else if iStr[j] < prev {
			// character less than previous
			return false
		}

		if j+1 == len(iStr) || iStr[j] != prev {
			// end of string, or character is diff to previous

			if accCheck(adjEqCount) {
				doubleFound = true
			}
			adjEqCount = 0
		}

		prev = iStr[j]
	}

	if doubleFound {
		return true
	}
	return false
}
