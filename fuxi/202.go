package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for {
		sum := getSum(n)
		if sum == 1 {
			return true
		}
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}
		n = sum
	}
}

func getSum(n int) int {
	sum := 0
	for n != 0 {
		temp := n % 10
		sum += temp * temp
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}
