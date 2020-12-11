package main

import (
	"reflect"
	"testing"
)

var eg1 = map[int][]int{
	0:  {1},
	1:  {4},
	4:  {5, 6, 7},
	5:  {6, 7},
	6:  {7},
	7:  {10},
	10: {11, 12},
	11: {12},
	12: {15},
	15: {16},
	16: {19},
	19: {22},
}

func TestBuildGraph(t *testing.T) {
	numbers := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	actual := buildAdjlist(numbers)

	if !reflect.DeepEqual(actual, eg1) {
		t.Errorf("got %+v, expected %+v", actual, eg1)
	}

}
