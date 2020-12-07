package main

import (
	"io/ioutil"
	"testing"
)

func TestAEg1(t *testing.T) {
	content, _ := ioutil.ReadFile("testdata/eg1.txt")
	expected := 4

	actual := run(string(content), false)

	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}

}

func TestBEg1(t *testing.T) {
	content, _ := ioutil.ReadFile("testdata/eg1.txt")
	expected := 32

	actual := run(string(content), true)

	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}
}

func TestBEg2(t *testing.T) {
	content, _ := ioutil.ReadFile("testdata/eg2.txt")
	expected := 126

	actual := run(string(content), true)

	if actual != expected {
		t.Errorf("Got %d, expected %d", actual, expected)
	}
}
