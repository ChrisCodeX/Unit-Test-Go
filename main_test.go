package main

import "testing"

// Function to test the addition function
func TestAddition(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Addition(item.a, item.b)
		if total != item.n {
			t.Errorf("Addition was incorrect, got %d expected %d", total, 10)
		}
	}
}

// Function to test the substraction function
func TestSubtraction(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{5, 4, 1},
		{6, 2, 4},
	}
	for _, item := range tables {
		s := Subtraction(item.a, item.b)
		if item.n != s {
			t.Errorf("Subtraction was incorrect, got %d, expected %d", s, item.n)
		}
	}
}

// Function to test the max function
func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}
}

// Function to test Fibonacci function
func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{44, 701408733},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}
