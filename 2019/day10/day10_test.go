package main

import (
	"math"
	"testing"
)

func TestCalculatingAngle0(t *testing.T) {
	c1 := coord{1, 2}
	c2 := coord{1, 1}
	angle := calcAngle(c1, c2)
	checkApproxEqual(float64(angle), float64(0), t)
}

func TestCalculatingAngleLessThan180(t *testing.T) {
	c1 := coord{1, 2}
	c2 := coord{2, 3}
	angle := calcAngle(c1, c2)

	checkApproxEqual(float64(angle), 2.356194, t)
}

func TestCalculatingAngle90(t *testing.T) {
	c1 := coord{1, 2}
	c2 := coord{3, 2}
	angle := calcAngle(c1, c2)
	checkApproxEqual(float64(angle), math.Pi/2, t)
}

func TestCalculatingAngle180(t *testing.T) {
	c1 := coord{1, 2}
	c2 := coord{1, 5}
	angle := calcAngle(c1, c2)
	checkApproxEqual(float64(angle), math.Pi, t)
}

func TestCalculatingAngleMoreThan180(t *testing.T) {
	c1 := coord{2, 2}
	c2 := coord{1, 5}
	angle := calcAngle(c1, c2)
	checkApproxEqual(float64(angle), 3.463300, t)
}

func TestSortCoords(t *testing.T) {
	sorted := sortCoordsByDist(coord{3, 3}, []coord{coord{9, 9}, coord{3, 3}, coord{3, 4}, coord{4, 4}})
	expected := []coord{coord{3, 3}, coord{3, 4}, coord{4, 4}, coord{9, 9}}

	checkEqual(sorted, expected, t)
}

func checkApproxEqual(actual float64, expected float64, t *testing.T) {
	if abs(actual-expected) > 0.00001 {
		t.Errorf("got %f, expected %f", actual, expected)
	}
}

func checkEqual(actual []coord, expected []coord, t *testing.T) {
	if len(actual) != len(expected) {
		t.Errorf("got %v, expected %v", actual, expected)
		return
	}

	for i, actualVal := range actual {
		expectedVal := expected[i]
		if actualVal != expectedVal {
			t.Errorf("Got %v, expected %v", actualVal, expectedVal)
		}
	}

}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
