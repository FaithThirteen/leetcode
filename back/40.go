package main

import (
	"fmt"
	"sort"
)

var result40 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	history := map[int]bool{}
	result40 = [][]int{}
	sort.Ints(candidates)
	backtracking40(len(candidates), 0, target, history, candidates, []int{})
	return result40
}

func backtracking40(n, start, target int, used map[int]bool, candidates, path []int) {
	sum := 0
	for _, v := range path {
		sum += v
	}
	if sum > target {
		return
	} else if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		result40 = append(result40, temp)
	}

	for i := start; i < n; i++ {
     	// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		// 因为排序过，前后相等值已经相等，map可以不用记录值，只需要记录索引，上一个的索引是false说明使用过了，同层的不可再使用
		// 这一步是去重的
		if i >0 && candidates[i] == candidates[i-1] && used[i-1] == false{
			continue
		}

		path = append(path, candidates[i])
		used[i] = true  // 树枝可以重复使用
		backtracking40(n, i+1, target, used, candidates, path)
		path = path[:len(path)-1]
		used[i] = false  // 回溯，同层不可以使用
	}

}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
