package main

import (
	"fmt"
	"sort"
)

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	if len(input) < 2 && k <= len(input){
		return input
	}
	if k < 1{
		return nil
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	var result = []int{input[0]}
	for i := 1; i < len(input); i++ {
		if len(result) == k{
			return result
		}
		result = append(result,input[i])
	}

	return result
}

func main() {

	fmt.Println(GetLeastNumbers_Solution([]int{234,233,233,233,233,233,233,233,233,233,233,233,233,233,233,233,233,233,233,233},10))
}