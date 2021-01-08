package main

import (
	"calculator"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var f func(a, b float64) float64
	// operator := os.Getenv("OPERATOR")
	// fmt.Println(operator)
	if len(os.Args) != 4 {
		log.Fatalf("bad args %v", os.Args)
	}
	a, b, c := os.Args[1], os.Args[2], os.Args[3]
	operator := c
	switch operator {
	case "add":
		f = calculator.Add
	case "subtract":
		f = calculator.Subtract
	default:
		log.Fatalf("bad operator %q", operator)
	}
	x, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatalf("bad input %q (expected number)", a)
	}
	y, err := strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatalf("bad input %q (expected number)", b)
	}
	answer := f(x, y)
	fmt.Println("The answer is", answer)
}
