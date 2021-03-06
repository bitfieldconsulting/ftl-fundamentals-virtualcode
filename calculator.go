// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the multiplied result
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers.  Returns 1st divided by 2nd or error
// checks for division by zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero undefined")
	}
	return a / b, nil
}

// SquareRoot of a number. Return square root or error
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of a number less than 0 is undefined")
	}
	return math.Sqrt(a), nil
}
