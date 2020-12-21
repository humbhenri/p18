package main

import (
	"testing"
)

func TestPrecedence(t *testing.T) {
	t1 := token{value: "+", token_type: add_t}
	t2 := token{value: "*", token_type: mul_t}
	if t1.precedence() <= t2.precedence() {
		t.Errorf("+ must have higher precedence than *, + precedence is %d, * precedence is %d",
			t1.precedence(), t2.precedence())
	}
}

func TestExpressions(t *testing.T) {
	expr := infixToPostfix("1 + 2 * 3 + 4 * 5 + 6")
	result, _ := eval(expr)
	if result != 231 {
		t.Errorf("Expected %d but was %d", 231, result)
	}
}

func TestParentheses(t *testing.T) {
	expr := infixToPostfix("1 + (2 * 3) + (4 * (5 + 6))")
	result, _ := eval(expr)
	if result != 51 {
		t.Errorf("Expected %d but was %d", 231, result)
	}
}
