package main

import "fmt"

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}

	backtracking(n, k, 1, []int{})
	return res
}

func backtracking(n, k, startIndex int, track []int) {
	if len(track) == k {
		temp := make([]int,k)
		copy(temp,track)
		res = append(res, temp)
	}

	if len(track)+n-startIndex+1 < k { // 假如从start开始后的元素已经加上track已经不满足于k个数的需求，可以忽略，因为是无效遍历
		return
	}

	for i := startIndex; i <= n; i++ {
		track = append(track, i)
		backtracking(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(combine(1, 1))
}
