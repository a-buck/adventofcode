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
	"strings"
)

var (
	inputFilePath = flag.String("input", "day04.txt", "input file path")
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

	valid := 0

	for {
		//passport

		fields := make(map[string]string)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			kvs := strings.Split(line, " ")
			for _, t := range kvs {
				parts := strings.Split(t, ":")
				k := parts[0]
				fields[k] = parts[1]
			}
		}

		if len(fields) == 0 {
			// processed all passports
			break
		}

		if containsAll(fields, "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid") {
			if partB {
				if hasValidValues(fields) {
					valid++
				}

			} else {
				// part A
				valid++
			}
		}
	}

	return valid
}

func validate(key, val string) bool {
	switch key {
	case "byr":
		return betweenInc(val, 1920, 2002)
	case "iyr":
		return betweenInc(val, 2010, 2020)
	case "eyr":
		return betweenInc(val, 2020, 2030)
	case "hgt":
		expr := regexp.MustCompile(`^(\d+)([a-z]+)$`)
		match := expr.FindStringSubmatch(val)
		if len(match) != 3 {
			return false
		}
		heightUnits := match[2]
		if heightUnits == "cm" {
			return betweenInc(match[1], 150, 193)
		} else if heightUnits == "in" {
			return betweenInc(match[1], 59, 76)
		} else {
			// units not recognised
			return false
		}
	case "hcl":
		expr := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		match := expr.FindStringSubmatch(val)
		return len(match) > 0
	case "ecl":
		return oneof(val, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
	case "pid":
		expr := regexp.MustCompile(`^\d{9}$`)
		match := expr.FindStringSubmatch(val)
		return len(match) > 0
	case "cid":
		return true
	}
	// field not recognised
	return false
}

func hasValidValues(fields map[string]string) bool {
	for k, v := range fields {
		if !validate(k, v) {
			return false
		}
	}
	return true
}

func oneof(val string, s []string) bool {
	for _, t := range s {
		if val == t {
			return true
		}
	}
	return false
}

func containsAll(m map[string]string, keys ...string) bool {
	for _, f := range keys {
		_, ok := m[f]
		if !ok {
			return false
		}
	}
	return true
}

func betweenInc(v string, lo, hi int) bool {
	i, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return i >= lo && i <= hi
}
