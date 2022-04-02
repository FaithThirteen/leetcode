package main

import "fmt"

// 我这里的策略是只要下一个能够跳的选择范围超过上一个，就选择下一个
func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	var end = len(nums)
	var step int
	for i := 0; i < end; i++ {
		if nums[i] == 0 && step <= i{ // 遇到0，如果当前最大的步数超过不了这个索引，则失败
			return false
		}
		if nums[i]+i > step {  // 如果可以比上一次走的更远，更新可以走的范围
			step = nums[i] + i
		}
		if step >= end-1 { // 如果可以走的范围大过数组索引，则可以到达
			return true
		}
	}

	return false
}

func main() {

	fmt.Println(canJump([]int{3,2,1,0,4}))
}
