package main

import "testing"

func TestAEg1(t *testing.T) {
	actual := getSeatID("FBFBBFFRLR")
	expected := 357

	if actual != expected {
		t.Errorf("Got %+v, expected %+v", actual, expected)
	}
}
