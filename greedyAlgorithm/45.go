package main

import "fmt"

func jump(nums []int) int {
	var steps, end, maxPosition int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > maxPosition {
			maxPosition = nums[i] + i
		}
		if i == end { // 走到了第一步覆盖范围尽头，如果还没结束，则取了范围内最大的值作为下一个最大范围
			end = maxPosition
			steps++
		}
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4,2,1}))
}
