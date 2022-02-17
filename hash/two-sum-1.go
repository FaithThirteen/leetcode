package main

import "fmt"

func main() {
	nums := []int{3,2,4}
	target := 6
	fmt.Println(twoSum(nums,target))
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
