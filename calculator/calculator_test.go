package calculator_test

import (
	"calculator"
	"math"
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
		{a: 1, b: -1, want: 0},
		{a: -3, b: 2, want: -1},
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
		{a: 4, b: 2, want: 2},
		{a: 1, b: -1, want: 2},
		{a: -3, b: 2, want: -5},
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: -1},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
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
		{a: 0, b: 5, want: 0},
		{a: -2, b: 3, want: -6},
		{a: -2, b: -3, want: 6},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 6, b: 3, want: 2},
		{a: -1, b: -1, want: 1},
		{a: -6, b: 3, want: -2},
		{a: 1, b: 3, want: 0.333333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Errorf("want error for division by zero, got nil")
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Errorf("want error for square root of negative number, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		input float64
		want  float64
	}
	testCases := []testCase{
		{input: 4, want: 2},
		{input: 9, want: 3},
		{input: 0, want: 0},
		{input: 1, want: 1},
		{input: 144, want: 12},
		{input: 144.36, want: 12.01},
		{input: 2, want: 1.41421356237},
		{input: 1.4, want: 1.2},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.input, tc.want, got)
		}
	}
}
