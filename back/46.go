package main

import "fmt"

var result46 [][]int

func permute(nums []int) [][]int {
	result46 = [][]int{}
	backtracking46(len(nums), len(nums), nums, []int{})
	return result46
}

func backtracking46(l, n int, nums, path []int) {
	if len(path) == l {
		temp := make([]int, len(path))
		copy(temp, path)
		result46 = append(result46, temp)
		return
	}

	for i := 0; i < n; i++ {
		path = append(path, nums[i])
		temp := make([]int, len(nums))
		copy(temp, nums)
		temp = append(temp[:i], temp[i+1:]...) // 除了自身，其余元素都可以选择，并且从开头选取

		backtracking46(l, len(temp), temp, path)
		path = path[:len(path)-1]
	}

}

func main() {

	fmt.Println(permute([]int{1, 2, 3}))
}
