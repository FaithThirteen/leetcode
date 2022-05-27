package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quick(x, n)
	}
	return 1.0 / quick(x, -n)
}
func quick(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quick(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	fmt.Println(myPow(2.10000, 3))
}
