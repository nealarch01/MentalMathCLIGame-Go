package main

import (
	"errors"
)

type ExprQueue struct {
	expressions []Expression // Expressions array will not be exported for abstraction purposes
}

// Methods

func (eq *ExprQueue) Count() int {
	return len(eq.expressions)
}

func (eq *ExprQueue) Pop() error {
	if eq.Count() == 0 {
		return errors.New("Queue underflow")
	}

	eq.expressions = eq.expressions[1:] // Removes the first index by slicing
	return nil
}

func (eq *ExprQueue) Push(expr Expression) {
	eq.expressions = append(eq.expressions, expr)
}

func (eq *ExprQueue) Top() (Expression, error) {
	if eq.Count() == 0 {
		var emptyExpr Expression
		return emptyExpr, errors.New("Queue underflow")
	}

	return eq.expressions[0], nil
}
