package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] = min746(cost[i-1], cost[i-2]) + cost[i] // 之所以加cost[i]，是因为支付了这个费用，可以向上爬了
	}
	fmt.Println(cost)
	// 注意最后一步可以理解为不用花费，所以取倒数第一步，第二步的最少值
	// 其实最后一步可以理解为，最后一次的时候已经支付了
	return min746(cost[len(cost)-1], cost[len(cost)-2])
}

// 方便理解
func minCostClimbingStairs1(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min746(dp[i-1], dp[i-2]) + cost[i]
	}

	return min746(dp[len(cost)-1], dp[len(cost)-2])
}

func min746(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1,100,1,1,1,100}))
}
