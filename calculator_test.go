package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
		{a: 100, b: 200, want: 300},
		{a: 0, b: -10, want: -10},
		{a: 0, b: 1.0, want: 1.00},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 12, b: 6, want: 6},
		{a: 0, b: 1, want: -1},
		{a: 1.1, b: 10, want: -8.9},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 3, want: 6},
		{a: 10, b: 10, want: 100},
		{a: 0, b: 1, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("multipy(%f x %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 1, b: 1, want: 1, errExpected: false},
		{a: 10, b: 5, want: 2, errExpected: false},
		{a: 10, b: 1, want: 10, errExpected: false},
		{a: 2, b: 0, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 4, want: 2, errExpected: false},
	}
	for _, tc := range testCases {
		got := calculator.SquareRoot(tc.a)
		if got != tc.want {
			t.Error("Sqrt incorrect")
		}
	}
}
