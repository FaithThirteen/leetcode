package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] { // 遇到重叠的区间，此时合并，因为更新了有边界，此时i不能前进
			intervals[i-1] = []int{intervals[i-1][0], max56(intervals[i][1], intervals[i-1][1])} // 左边界已经排序，更新最大的右边界
			intervals = append(intervals[:i], intervals[i+1:]...) // 因为合并，去掉了第i个区间，后面的往前移
			i--
		}
	}
	return intervals
}
func max56(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
