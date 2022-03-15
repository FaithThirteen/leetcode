package main

import "fmt"

var result491 [][]int

func findSubsequences(nums []int) [][]int {
	result491 = [][]int{}
	backtracking491(len(nums), 0, nums, []int{})
	return result491
}

func backtracking491(n, start int, nums, path []int) {

	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		result491 = append(result491, temp)
	}

	history := make(map[int]bool)
	for i := start; i < n; i++ {
		// 1.如果当前的元素小于path数组最后一个元素，跳过此元素
		// 2.同树层不可使用同一个元素
		if (len(path) > 0 && nums[i] < path[len(path)-1]) || history[nums[i]] {
			continue
		}

		path = append(path, nums[i])
		history[nums[i]] = true
		backtracking491(n, i+1, nums, path)
		path = path[:len(path)-1]

	}
}

func main() {

	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
