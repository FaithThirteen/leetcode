package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] { // nums[i]充当的是一个临时累加和变量
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {  // 收集累计和
			maxSum = nums[i]
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}