package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/a-buck/adventofcode/2020/stack"
	"github.com/a-buck/adventofcode/2020/utils"
)

var (
	inputFilePath = flag.String("input", "day18.txt", "input file path")
	partB         = flag.Bool("partB", false, "enable part b")
)

type token interface {
	string() string
}

type operator rune
type operand int

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, l := range strings.Split(string(content), "\n") {
		tokens := parseTokens(string(l))
		postfix := toPostfixExpr(tokens, *partB)
		res := evaluatePostfixExpr(postfix)
		sum += res
	}

	fmt.Println(sum)
}

func parseTokens(in string) []token {
	tokens := make([]token, 0)

	for _, r := range in {
		if r == ' ' {
			continue
		}

		if r >= '0' && r <= '9' {
			v := utils.ToInt(string(r))
			tokens = append(tokens, operand(v))
		} else {
			tokens = append(tokens, operator(r))
		}
	}
	return tokens
}

// implemented with variation of shunting yard algorithm
func toPostfixExpr(tokens []token, partB bool) []token {
	operatorPrecedence := map[operator]int{
		'*': 0,
		'+': 0,
	}
	if partB {
		// in part B, + has higher precedence
		operatorPrecedence['+'] = 1
	}

	operatorStack := make(stack.Stack, 0)
	out := make([]token, 0)

	for _, token := range tokens {

		if v, ok := token.(operand); ok {
			out = append(out, v)
		} else {
			v := token.(operator)

			if v.isBinaryOperator() {
				for {
					top, ok := operatorStack.Peek().(operator)
					if ok &&
						// there is an operator at the top of the operator stack
						top.isBinaryOperator() &&
						// the operator at the top has greater precedence
						operatorPrecedence[top] >= operatorPrecedence[v] {
						out = append(out, top)
						operatorStack.Pop()
					} else {
						break
					}
				}
				operatorStack = append(operatorStack, v)
			} else if v == '(' {
				operatorStack.Push(token)
			} else if v == ')' {
				for operatorStack.Peek() != nil && operatorStack.Peek().(operator) != '(' {
					out = append(out, operatorStack.Pop().(operator))
				}
				operatorStack.Pop() // remove '('
			}

		}

	}

	for i := len(operatorStack) - 1; i >= 0; i-- {
		out = append(out, operatorStack[i].(operator))
	}

	return out
}

func evaluatePostfixExpr(tokens []token) int {
	stack := stack.Stack{}

	for _, t := range tokens {
		if v, ok := t.(operand); ok {
			stack.Push(v)
		} else {
			// t is an operator
			v := t.(operator)

			operand1 := stack.Pop().(operand)
			operand2 := stack.Pop().(operand)
			stack.Push(operand(v.execute(operand1, operand2)))
		}
	}

	if len(stack) != 1 {
		log.Fatalf("Expected 1 item on stack at end, got %d", len(stack))
	}

	return int(stack[0].(operand))
}

func (o operand) string() string {
	return strconv.Itoa(int(o))
}

func (o operator) string() string {
	return string(o)
}

func (o operator) execute(a, b operand) int {
	if o == '*' {
		return int(a * b)
	}
	// must be +
	return int(a + b)
}

func (o operator) isBinaryOperator() bool {
	return o == '+' || o == '*'
}
