package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {

	// 按照x开始大小排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var result = 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] { //如果前一位的右边界小于后一位的左边界，则一定不重合
			result++
		} else { // 气球i和气球i-1挨着
			points[i][1] = min(points[i-1][1], points[i][1]) // 更新重叠气球最小右边界,覆盖该位置的值，留到下一步使用
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	fmt.Println(findMinArrowShots([][]int{{2, 3}, {2, 3}}))
}
