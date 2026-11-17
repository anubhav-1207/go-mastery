// Package main implements a small CLI calculator.
package main

import (
	"errors"
	"fmt"
	"math"
)

// ErrDivisionByZero is returned when dividing by zero.
var ErrDivisionByZero = errors.New("division by zero")

// Add returns the sum of a and b.
func Add(a, b float64) float64 { return a + b }

// Subtract returns a minus b.
func Subtract(a, b float64) float64 { return a - b }

// Multiply returns a times b.
func Multiply(a, b float64) float64 { return a * b }

// Divide returns a divided by b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Power returns a raised to b.
func Power(a, b float64) float64 { return math.Pow(a, b) }

// Calculate dispatches to the named operation.
func Calculate(op string, a, b float64) (float64, error) {
	switch op {
	case "add", "+":
		return Add(a, b), nil
	case "sub", "-":
		return Subtract(a, b), nil
	case "mul", "*":
		return Multiply(a, b), nil
	case "div", "/":
		return Divide(a, b)
	case "pow", "^":
		return Power(a, b), nil
	default:
		return 0, fmt.Errorf("unknown operation %q", op)
	}
}
