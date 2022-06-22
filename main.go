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

func SumFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return SumFibonacci(n-1) + SumFibonacci(n-2)
}
