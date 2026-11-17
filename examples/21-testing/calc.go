// Package calc provides simple arithmetic used to demonstrate testing.
package calc

import "errors"

// ErrDivisionByZero is returned by Divide when the divisor is zero.
var ErrDivisionByZero = errors.New("division by zero")

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Divide returns a/b and an error if b is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Reverse returns its input reversed rune-by-rune.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
