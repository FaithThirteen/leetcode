package main

import (
	"fmt"
	"sort"
)

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	if len(numbers) < 2 {
		return numbers[0]
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	return numbers[len(numbers)/2]
}

func main() {
	//fmt.Println(MoreThanHalfNum_Solution([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(MoreThanHalfNum_Solution2([]int{7, 3, 4, 5, 2, 7}))
}

func MoreThanHalfNum_Solution2(nums []int) int {
	var candidate, count = nums[0], 1
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = nums[i+1]
			count++
		}
		fmt.Printf("candidate： %d, count: %d \n",candidate,count)
	}
	return candidate
}
