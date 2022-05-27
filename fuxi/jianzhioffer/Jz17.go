package main

import (
	"fmt"
	"math"
)

func printNumbers( n int ) []int {
	// write code here
	number := math.Pow(10,float64(n))
	num := int(number) -1
	res := make([]int,num)
	for i:=1;i<=num;i++{
		res[i-1] = i
	}
	return res
}

func main() {
	fmt.Println(printNumbers(1))
}