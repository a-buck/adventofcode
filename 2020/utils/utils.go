package utils

import (
	"log"
	"regexp"
	"strconv"
)

// Match regex expr on string
func Match(s string, expr *regexp.Regexp) map[string]string {
	match := expr.FindStringSubmatch(s)
	result := make(map[string]string)
	for i, name := range expr.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

// ToUint converts a string to a uint64
func ToUint(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		log.Fatalf("unable to parse %s to int", s)
	}
	return v
}
