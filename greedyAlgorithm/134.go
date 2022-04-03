package main

import "fmt"

// i从0开始累加rest[i]，和记为curSum，一旦curSum小于零，说明[0, i]区间都不能作为起始位置，起始位置从i+1算起，再从0计算curSum。
//那么为什么一旦[i，j] 区间和为负数，起始位置就可以是j+1呢，j+1后面就不会出现更大的负数？
//
//如果出现更大的负数，就是更新j，那么起始位置又变成新的j+1了。
//
//而且j之前出现了多少负数，j后面就会出现多少正数，因为耗油总和是大于零的（前提我们已经确定了一定可以跑完全程）。
//
//那么局部最优：当前累加rest[j]的和curSum一旦小于0，起始位置至少要是j+1，因为从j开始一定不行。全局最优：找到可以跑一圈的起始位置。
func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 { // 总和少于0，必定不可能到达终点
		return -1
	}
	return start
}

func main() {
	//fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
