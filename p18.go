package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

// shunting yard algorithm
func infixToPostfix(input string) []token {
	var output []token
	stack := stack.New()
	for _, t := range tokenize(strings.NewReader(input)) {
		switch t.token_type {
		case number_t:
			output = append(output, t)
		case open_paren_t:
			stack.Push(t)
		case close_paren_t:
			for stack.Len() != 0 && stack.Peek().(token).token_type != open_paren_t {
				output = append(output, stack.Pop().(token))
			}
			stack.Pop()
		default:
			for stack.Len() != 0 && stack.Peek().(token).isOperator() && t.precedence() <= stack.Peek().(token).precedence() {
				output = append(output, stack.Pop().(token))
			}
			stack.Push(t)
		}
	}
	for stack.Len() != 0 {
		output = append(output, stack.Pop().(token))
	}
	return output
}

func eval(infix []token) (int, error) {
	stack := stack.New()
	for _, t := range infix {
		switch t.token_type {
		case number_t:
			stack.Push(t.value.(int))
		case add_t:
			operand1 := stack.Pop().(int)
			operand2 := stack.Pop().(int)
			stack.Push(operand1 + operand2)
		case sub_t:
			operand1 := stack.Pop().(int)
			operand2 := stack.Pop().(int)
			stack.Push(operand1 - operand2)
		case mul_t:
			operand1 := stack.Pop().(int)
			operand2 := stack.Pop().(int)
			stack.Push(operand1 * operand2)
		default:
			return 0, fmt.Errorf("Invalid argument: %q", t)
		}
	}
	return stack.Pop().(int), nil
}

func main() {
	file, err := os.Open("./i18")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		expr := infixToPostfix(scanner.Text())
		result, err := eval(expr)
		if err != nil {
			log.Fatal(err)
		}
		total += result
	}
	fmt.Println(total)
}
