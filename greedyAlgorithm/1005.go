package main

import (
	"fmt"
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	i := len(nums) - 1
	for k != 0 {
		nums[i] *= -1
		k--
	}
	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

func main() {

	fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
}
