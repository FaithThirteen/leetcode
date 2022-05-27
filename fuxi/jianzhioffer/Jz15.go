package main

import "fmt"

func NumberOf1(n int) int {
	// write code here
	var res int
	for n != 0 {
		res += n & 1
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(NumberOf1(11))
}

func hammingWeight(num uint32) int {
	// write code here
	var res uint32
	for num != 0 {
		res += num & 1
		num >>= 1
	}
	return int(res)
}