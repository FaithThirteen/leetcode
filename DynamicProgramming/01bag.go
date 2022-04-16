package main

import "fmt"

func bagMaxValue(bagWeight int, weight, value []int) int {
	dp := make([][]int, len(weight))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bagWeight+1, bagWeight+1)
	}

	// 初始化
	for j := bagWeight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}

	for i := 1; i < len(weight); i++ { // 遍历物品
		for j := weight[i]; j <= bagWeight; j++ { // 遍历背包
			dp[i][j] = max01bag(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}
	return dp[len(weight)-1][bagWeight]
}

func max01bag(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(bagMaxValue(4, []int{1, 3, 4}, []int{15, 20, 30}))
}
