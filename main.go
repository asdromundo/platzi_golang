package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return (Fibonacci(n-1) + Fibonacci(n-2))
}

func UseTailFibonacci(n int) int {
	return TailFibonacci(n, 0, 1)
}

func TailFibonacci(n, a, b int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else {
		return TailFibonacci(n-1, b, a+b)
	}
}
