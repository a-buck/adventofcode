package main

import (
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	r := strings.NewReader("10 ORE => 10 A")
	equations := parseInput(r)

	checkIntEq(len(equations), 1, t)

	eqn := equations[0]
	checkIntEq(len(eqn.inputs), 1, t)

	checkIntEq(eqn.inputs[0].amount, 10, t)
	checkStrEq(eqn.inputs[0].name, "ORE", t)

	checkIntEq(eqn.output.amount, 10, t)
	checkStrEq(eqn.output.name, "A", t)

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
