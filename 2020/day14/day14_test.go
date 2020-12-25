package main

import (
	"reflect"
	"testing"
)

func TestGenerateMaskOptions(t *testing.T) {

	options := maskOptions("X1001X")

	expected := []string{"000000", "000001", "100000", "100001"}

	if !reflect.DeepEqual(options, expected) {
		t.Errorf("got %+v, expected %+v", options, expected)
	}
}
