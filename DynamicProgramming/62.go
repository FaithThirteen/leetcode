package main

import "fmt"

func uniquePaths(m int, n int) int {

	if m == 1 || n == 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 贴边的两个只有一种方法可以到达，因为只可以向下和向右移动
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			// 每个格子等于它上面的格子加左边的格子的路径和，dp[i][j] = dp[i-1][j] + dp[i][j-1]
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 3))
}
