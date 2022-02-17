package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去掉重复的三元组
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && n2 == nums[l] {
					l++
				}
				for l < r && n3 == nums[r] {
					r--
				}
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			}
		}

	}
	return res
}
