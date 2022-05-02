package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {

	sum := 0
	l, r := 0, 0

	result := len(nums)+1

	for r < len(nums) {
		sum += nums[r]
		for sum >= target{
			sub := r-l+1
			if sub < result{
				result = sub
			}
			sum -= nums[l]
			l++
		}
		r++
	}

	if result == len(nums)+1{
		return 0
	}

	return result
}

func main() {

	fmt.Println(minSubArrayLen(7,[]int{2,3,1,2,4,3}))
}