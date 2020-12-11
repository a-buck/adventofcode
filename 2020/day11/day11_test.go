package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNext1PartB(t *testing.T) {
	actual := toGrid(open("testdata/eg1B0.txt", t))
	expected := toGrid(open("testdata/eg1B2.txt", t))

	newGrid, _ := actual.next(true)

	if !reflect.DeepEqual(newGrid, expected) {
		t.Errorf("\nactual: \n%s\nexpected: \n%s", newGrid.toString(), expected.toString())
	}
}

func TestNext2PartB(t *testing.T) {
	actual := toGrid(open("testdata/eg1B2.txt", t))
	expected := toGrid(open("testdata/eg1B3.txt", t))

	newGrid, _ := actual.next(true)

	if !reflect.DeepEqual(newGrid, expected) {
		t.Errorf("\nactual: \n%s\nexpected: \n%s", newGrid.toString(), expected.toString())
	}
}

func open(filePath string, t *testing.T) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	return content
}
