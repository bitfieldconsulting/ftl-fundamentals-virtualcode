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
	operator := os.Getenv("OPERATOR")
	switch operator {
	case "add":
		f = calculator.Add
	case "subtract":
		f = calculator.Subtract
	default:
		log.Fatalf("bad operator %q", operator)
	}
	if len(os.Args) != 3 {
		log.Fatalf("bad args %v", os.Args)
	}
	a, b := os.Args[1], os.Args[2]
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