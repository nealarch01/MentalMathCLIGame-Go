package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Expression struct {
	leftOperand  int
	rightOperand int
	operator     string
}

func (e *Expression) Init() {
	rand.Seed(time.Now().UnixNano())
	operators := []string{"+", "-", "*", "/"}

	e.operator = operators[rand.Intn(len(operators))]

	var min int
	var max int

	if e.operator == "+" || e.operator == "-" { // Addition / subtraction
		max = 200
		min = 1
	} else if e.operator == "*" { // Multiplication
		max = 12
		min = 2
	} else { // Division
		max = 100
		min = 2
	}

	e.leftOperand = rand.Intn((max - min) + min)
	e.rightOperand = rand.Intn((max - min) + min)

	// If the left operand is smaller, swap with the right
	// ie 2 / 5 = 0 (not fun)
	// 5 / 2 = 2 (more fun)!
	if (e.operator == "-" || e.operator == "/") && e.leftOperand < e.rightOperand {
		// Do a swap
		temp := e.leftOperand
		e.leftOperand = e.rightOperand
		e.rightOperand = temp
	}

	// Since this is passed by reference, no need to return Expression type
	// Properties are now initialized
}

func (e *Expression) Display() string {
	exprStr := fmt.Sprintf("%d %s %d", e.leftOperand, e.operator, e.rightOperand)
	return exprStr
}

func (e *Expression) CalcResult() int {
	switch e.operator {
	case "+":
		return e.leftOperand + e.rightOperand

	case "-":
		return e.leftOperand - e.rightOperand

	case "*":
		return e.leftOperand * e.rightOperand

	case "/":
		return e.leftOperand / e.rightOperand
	}

	return 0
}
