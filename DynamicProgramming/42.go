package main

import "fmt"

func trap(height []int) int {

	max_left := make([]int, len(height))  // max_lef[i]代表第i个的左边最大值
	max_right := make([]int, len(height)) // max_right[i] 代表第i个的右边最大值

	max_left[0] = height[0]
	max_right[0] = height[0]

	for i := 1; i < len(height); i++ {
		max_left[i] = max42(max_left[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		max_right[i] = max42(max_right[i+1], height[i+1])
	}

	var sum int
	for i := 0; i < len(height); i++ {
		min := min42(max_left[i], max_right[i]) // 找出短板
		if min > height[i] { // 左边或者右边最大值的最小值 大于我，说明可以承雨水
			sum += min - height[i]
		}
	}

	return sum
}

func max42(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min42(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
