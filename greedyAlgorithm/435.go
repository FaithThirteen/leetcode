package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result = 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] { // 有交集
			intervals[i][1] = min435(intervals[i][1], intervals[i-1][1])
			result++
		}
	}
	return result
}

func min435(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}
