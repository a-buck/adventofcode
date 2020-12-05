package main

import "testing"

func TestAEg1(t *testing.T) {
	actual := decodeSeat("FBFBBFFRLR")
	expected := seat{44, 5}

	if actual != expected {
		t.Errorf("Got %+v, expected %+v", actual, expected)
	}
}
