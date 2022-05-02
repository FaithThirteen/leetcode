package main

import (
	"fmt"
	"sort"
)

var result [][]int

// 用一下回溯
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	backtracking(0, len(nums), 0, nums, []int{}, map[int]bool{})
	return result
}

func backtracking(start, end, sum int, nums, path []int, used map[int]bool) {
	if len(path) == 3 {
		if sum != 0 {
			return
		}
		temp	 := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		return
	}

	for i := start; i < end; i++ {

		if i > 0 && i != 0 && nums[i] == nums[i-1] && used[i-1] == true {
			continue
		}
		sum += nums[i]
		path = append(path, nums[i])
		used[i] = false
		backtracking(i+1, end, sum, nums, path, used)
		path = path[:len(path)-1]
		used[i] = true
		sum -= nums[i]
	}
}

func main() {
	fmt.Println(threeSum([]int{34,55,79,28,46,33,2,48,31,-3,84,71,52,-3,93,15,21,-43,57,-6,86,56,94,74,83,-14,28,-66,46,-49,62,-11,43,65,77,12,47,61,26,1,13,29,55,-82,76,26,15,-29,36,-29,10,-70,69,17,49}))
}
