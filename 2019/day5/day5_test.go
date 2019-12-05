package main

import "testing"

func TestEchoProgram(t *testing.T) {

	prog := []string{"3", "0", "4", "0", "99"}

	doTest(prog, 7, 7, t)
	doTest(prog, 8, 8, t)

}

func TestEq8PosModeProgram(t *testing.T) {
	prog := []string{"3", "9", "8", "9", "10", "9", "4", "9", "99", "-1", "8"}

	doTest(prog, 7, 0, t)
	doTest(prog, 8, 1, t)
	doTest(prog, 9, 0, t)
}

func TestLessThan8PosModeProgram(t *testing.T) {
	prog := []string{"3", "9", "7", "9", "10", "9", "4", "9", "99", "-1", "8"}

	doTest(prog, 7, 1, t)
	doTest(prog, 8, 0, t)
	doTest(prog, 9, 0, t)
}

func TestEq8ImmModeProgram(t *testing.T) {
	prog := []string{"3", "3", "1108", "-1", "8", "3", "4", "3", "99"}

	doTest(prog, 7, 0, t)
	doTest(prog, 8, 1, t)
	doTest(prog, 9, 0, t)
}

func TestLessThan8ImmModeProgram(t *testing.T) {
	prog := []string{"3", "3", "1107", "-1", "8", "3", "4", "3", "99"}

	doTest(prog, 7, 1, t)
	doTest(prog, 8, 0, t)
	doTest(prog, 9, 0, t)
}

func TestJumpPosMode(t *testing.T) {
	prog := []string{"3", "12", "6", "12", "15", "1", "13", "14", "13", "4", "13", "99", "-1", "0", "1", "9"}

	doTest(prog, 0, 0, t)
	doTest(prog, 1, 1, t)
	doTest(prog, 2, 1, t)
}

func TestJumpImmMode(t *testing.T) {
	prog := []string{"3", "3", "1105", "-1", "9", "1101", "0", "0", "12", "4", "12", "99", "1"}

	doTest(prog, 0, 0, t)
	doTest(prog, 1, 1, t)
	doTest(prog, 2, 1, t)
}

func doTest(program []string, input int, expected int, t *testing.T) {
	output := run(program, input)

	if len(output) == 0 {
		t.Errorf("zero length output")
	}

	if output[len(output)-1] != expected {
		t.Errorf("got %d, wanted %d", output[len(output)-1], expected)
	}
}
