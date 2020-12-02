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
	inputFilePath = flag.String("input", "day02.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

func main() {
	flag.Parse()

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := run(file, *partB)

	fmt.Println(ans)
}

func run(r io.Reader, partB bool) int {

	scanner := bufio.NewScanner(r)

	validCount := 0

	expr := regexp.MustCompile(`(?P<left>\d+)-(?P<right>\d+)\s(?P<char>\w):\s(?P<password>\w+)`)

	for scanner.Scan() {
		line := scanner.Text()

		match := expr.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range expr.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		left, _ := strconv.Atoi(result["left"])
		right, _ := strconv.Atoi(result["right"])
		targetChar := []rune(result["char"])[0]
		password := []rune(result["password"])

		if partB {
			leftValid := password[left-1] == targetChar
			rightValid := password[right-1] == targetChar

			if (leftValid || rightValid) && !(leftValid && rightValid) {
				validCount++
			}

		} else {
			// part A
			charCount := 0
			for _, c := range password {
				if c == targetChar {
					charCount++
				}
			}

			if charCount >= left && charCount <= right {
				validCount++
			}
		}

	}

	return validCount

}
