package main

import (
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	r := strings.NewReader("10 ORE => 10 A")
	graph := parseInput(r)

	v := graph["A"]

	checkIntEq(v.amountGeneratedByReac, 10, t)
	checkIntEq(len(v.edges), 1, t)

	e := v.edges[0]

	checkIntEq(e.amount, 10, t)
	checkStrEq(e.dest, "ORE", t)

}

func TestMultipleInputs(t *testing.T) {
	r := strings.NewReader("11 QMDTJ, 15 LVPK, 5 LZPCS => 3 KJVZ")
	graph := parseInput(r)

	v := graph["KJVZ"]

	checkIntEq(v.amountGeneratedByReac, 3, t)
	checkIntEq(len(v.edges), 3, t)

	checkStrEq(v.edges[0].dest, "QMDTJ", t)
	checkStrEq(v.edges[1].dest, "LVPK", t)
	checkStrEq(v.edges[2].dest, "LZPCS", t)

	checkIntEq(v.edges[0].amount, 11, t)
	checkIntEq(v.edges[1].amount, 15, t)
	checkIntEq(v.edges[2].amount, 5, t)

}

func checkIntEq(act, exp int, t *testing.T) {
	if act != exp {
		t.Errorf("got %d, expected %d", act, exp)
	}
}

func checkStrEq(act, exp string, t *testing.T) {
	if act != exp {
		t.Errorf("got %s, expected %s", act, exp)
	}
}
