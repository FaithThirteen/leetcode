package main

import "fmt"

// 递归，效率低
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

func main() {
	fmt.Println(fib(4))
}
