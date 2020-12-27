package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/a-buck/adventofcode/2020/utils"
)

var (
	inputFilePath = flag.String("input", "day16.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
	ruleExpr      = regexp.MustCompile(`^(?P<key>[\w\s]+):\s(?P<leftLo>\d+)-(?P<leftHi>\d+)\sor\s(?P<rightLo>\d+)-(?P<rightHi>\d+)$`)
)

type rangeRule struct {
	lo, hi int
}

type ticket []int

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(content), "\n\n")

	rules := parseRules(parts[0])

	allNearByTickets := parseTickets(parts[2])

	validNearbyTickets, errorRate := getValidTickets(allNearByTickets, rules)

	// part A
	fmt.Println(errorRate)

	colToCandidates := createColToCandidates(validNearbyTickets, rules)

	correctPositions := createCorrectPositions(colToCandidates)

	myTicket := parseTickets(parts[1])[0]

	partBResult := 1
	for name, i := range correctPositions {
		if strings.HasPrefix(name, "departure") {
			partBResult *= myTicket[i]
		}
	}

	// part B
	fmt.Println(partBResult)
}

// returns (valid tickets, ticket scanning error rate)
func getValidTickets(tickets []ticket, rules map[string][]rangeRule) ([]ticket, int) {
	sum := 0
	validTickets := make([]ticket, 0)
	for _, ticket := range tickets {
		isValidTicket := true
		for _, val := range ticket {
			if !validateField(val, rules) {
				isValidTicket = false
				sum += val
			}
		}

		if isValidTicket {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, sum
}

func validateField(val int, rules map[string][]rangeRule) bool {
	for _, fieldRules := range rules {
		for _, rule := range fieldRules {
			if val >= rule.lo && val <= rule.hi {
				return true
			}
		}
	}
	return false
}

func parseRules(part string) map[string][]rangeRule {
	ruleLines := strings.Split(part, "\n")

	rules := make(map[string][]rangeRule)

	for _, v := range ruleLines {
		match := utils.Match(v, ruleExpr)
		r1 := rangeRule{lo: utils.ToInt(match["leftLo"]), hi: utils.ToInt(match["leftHi"])}
		r2 := rangeRule{lo: utils.ToInt(match["rightLo"]), hi: utils.ToInt(match["rightHi"])}
		rules[match["key"]] = []rangeRule{r1, r2}
	}
	return rules
}

func parseTickets(part string) []ticket {
	lines := strings.Split(part, "\n")[1:] // skip heading
	tickets := make([]ticket, len(lines))
	for i, ticketLine := range lines {
		ticket := toTicket(ticketLine)
		tickets[i] = ticket
	}
	return tickets
}

func createCorrectPositions(colToCandidates map[int][]string) map[string]int {
	seen := make(map[string]bool)
	correctPositions := make(map[string]int)

	changed := true
	for changed {
		changed = false
		for col, candidates := range colToCandidates {

			notseencandidates := make([]string, 0)
			for _, c := range candidates {
				if _, ok := seen[c]; !ok {
					notseencandidates = append(notseencandidates, c)
				}
			}

			if len(notseencandidates) == 1 {
				correctPositions[notseencandidates[0]] = col
				seen[notseencandidates[0]] = true
				changed = true
			}
		}
	}
	return correctPositions
}

func createColToCandidates(tickets []ticket, rules map[string][]rangeRule) map[int][]string {
	colToCandidates := make(map[int]map[string]bool)

	// row 0
	for i, v := range tickets[0] {
		cands := getCandidates(v, rules)
		colToCandidates[i] = cands
	}

	for col := 0; col < len(tickets[0]); col++ {
		for row := 1; row < len(tickets); row++ {
			val := tickets[row][col]
			newCandidates := getCandidates(val, rules)

			currentCandidates := colToCandidates[col]

			for c := range currentCandidates {
				if _, exists := newCandidates[c]; !exists {
					delete(currentCandidates, c)
				}
			}
		}
	}

	colToCandidates2 := make(map[int][]string)

	for k, v := range colToCandidates {
		lst := make([]string, 0)
		for k2 := range v {
			lst = append(lst, k2)
		}
		colToCandidates2[k] = lst
	}

	return colToCandidates2
}

func getCandidates(val int, keyToRules map[string][]rangeRule) map[string]bool {
	candidates := make(map[string]bool, 0)

	for k, v := range keyToRules {
		for _, r := range v {
			if r.validate(val) {
				candidates[k] = true
			}
		}
	}
	return candidates
}

func toTicket(s string) ticket {
	parts := strings.Split(s, ",")
	ticket := make([]int, len(parts))
	for i, v := range parts {
		ticket[i] = utils.ToInt(v)
	}
	return ticket
}

func (r *rangeRule) validate(val int) bool {
	return val >= r.lo && val <= r.hi
}
