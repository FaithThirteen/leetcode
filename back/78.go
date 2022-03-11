package main

import "fmt"

var result78 [][]int

func subsets(nums []int) [][]int {
	result78 = [][]int{[]int{}}
	backtracking78(len(nums),0,nums,[]int{})
	return result78
}

func backtracking78(n, start int, nums, path []int) {

	for i := start; i < n; i++ {
		path = append(path, nums[i])
		temp := make([]int, len(path))
		copy(temp, path)
		result78 = append(result78, temp)
		backtracking78(n, i+1, nums, path)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}