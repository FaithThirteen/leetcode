package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3,-2,-1,0,0,1,2,3}
	target := 0

	fmt.Println(fourSum(nums, target))

}

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	res := [][]int{}

	// 确定nums[i]和nums[k]的值
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		for k := i + 1; k < len(nums)-2; k++ {
			n2 := nums[k]
			if k > i+1 && n2 == nums[k-1] {
				continue
			}

			l, r := k+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && nums[l+1] == n3 {
						l++
					}
					for l < r && nums[r-1] == n4 {
						r--
					}
					l++
					r--
				} else if n1+n2+n3+n4 < target {
					l++
				} else if n1+n2+n3+n4 > target {
					r--
				}
			}

		}

	}

	return res
}
