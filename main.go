package main

func Addition(x, y int) int {
	return x + y
}

func Subtraction(x, y int) int {
	return x - y
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
