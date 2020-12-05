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

type validateFn func(string) bool

var (
	inputFilePath = flag.String("input", "day04.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
	validationFns = map[string]validateFn{
		"byr": validateByr,
		"iyr": validateIyr,
		"eyr": validateEyr,
		"hgt": validateHgt,
		"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString,
		"ecl": regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$").MatchString,
		"pid": regexp.MustCompile(`^\d{9}$`).MatchString,
		"cid": validateCid,
	}
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

func hasValidValues(fields map[string]string) bool {
	for k, v := range fields {
		if !validationFns[k](v) {
			return false
		}
	}
	return true
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

func validateByr(val string) bool {
	return betweenInc(val, 1920, 2002)
}

func validateIyr(val string) bool {
	return betweenInc(val, 2010, 2020)
}

func validateEyr(val string) bool {
	return betweenInc(val, 2020, 2030)
}

func validateHgt(val string) bool {
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
}

func validateCid(_ string) bool {
	return true
}
