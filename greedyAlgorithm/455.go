package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var result int
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			result++
			i++
			j++
		}else{
			j++
		}
	}
	return result
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1,1}))
}


//排序后，局部最优
func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	// 从小到大
	child := 0
	for sIdx := 0; child < len(g) && sIdx < len(s); sIdx++ {
		if s[sIdx] >= g[child] {//如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
			child++
		}
	}

	return child
}