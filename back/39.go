package main

import "fmt"

var result39 [][]int

func combinationSum(candidates []int, target int) [][]int {
	result39 = [][]int{}

	backtracking39(len(candidates),target,0,candidates,[]int{})

	return result39
}

func backtracking39(n, target ,start int, candidates, path []int) {
	sum := 0
	for _, v := range path {
		sum += v
	}
	if sum > target {
		return
	} else if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		result39 = append(result39, temp)
		return
	}


	for i := start; i < n; i++ {
		path = append(path, candidates[i])
		backtracking39(n,target, i, candidates, path)
		path = path[:len(path)-1]
	}

}

func main() {
	fmt.Println(combinationSum([]int{2,3,5},8))

}