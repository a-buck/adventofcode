package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/a-buck/adventofcode/2020/testutils"
)

func TestParseTokens(t *testing.T) {
	in := "123*+"

	out := parseTokens(in)

	expected := []token{operand(1), operand(2), operand(3), operator('*'), operator('+')}

	CheckEq(out, expected, t)

}

func TestToPostFix1(t *testing.T) {

	in := "1 + 2 * 3 + 4 * 5 + 6"

	expected := "12+3*4+5*6+"

	actual := toPostfixExpr(parseTokens(in), false)
	actualStr := make([]string, 0)

	for _, v := range actual {
		actualStr = append(actualStr, v.string())
	}

	testutils.CheckStringsEq(strings.Join(actualStr, ""), expected, t)
}

func TestToPostFixWithParentheses(t *testing.T) {
	in := "1 + (2 * 3)"

	expected := "123*+"

	actual := toPostfixExpr(parseTokens(in), false)

	actualStr := make([]string, 0)

	for _, v := range actual {
		actualStr = append(actualStr, v.string())
	}

	testutils.CheckStringsEq(strings.Join(actualStr, ""), expected, t)
}

func TestToPostFixWithParentheses2(t *testing.T) {
	in := "1 + (2 * 3) + (4 * (5 + 6))"

	expected := "123*+456+*+"

	actual := toPostfixExpr(parseTokens(in), false)

	actualStr := make([]string, 0)

	for _, v := range actual {
		actualStr = append(actualStr, v.string())
	}

	testutils.CheckStringsEq(strings.Join(actualStr, ""), expected, t)
}

func TestEvaluatePostfixExpr(t *testing.T) {

	in := []token{operand(1), operand(2), operand(3), operator('*'), operator('+')}

	expected := 7

	actual := evaluatePostfixExpr(in)

	testutils.CheckIntsEq(actual, expected, t)
}

func TestEgPartA(t *testing.T) {
	in := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"

	actual := evaluatePostfixExpr(toPostfixExpr(parseTokens(in), false))

	testutils.CheckIntsEq(actual, 13632, t)
}

func TestEgPartB(t *testing.T) {
	in := "2 * 3 + (4 * 5)"

	actual := evaluatePostfixExpr(toPostfixExpr(parseTokens(in), true))

	testutils.CheckIntsEq(actual, 46, t)
}

func CheckEq(actual, expected []token, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}
