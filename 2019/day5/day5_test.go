package main

import "testing"

func TestEchoProgram(t *testing.T) {

	input := 7

	answer, err := runProgram([]string{"3", "0", "4", "0", "99"}, true, input)

	if err != nil {
		t.Error(err)
	}

	// for the above program, expect input to be output as is.
	if answer != input {
		t.Errorf("got %d, wanted %d", answer, input)
	}

}
