package main

func main() {

}

// 动态规划
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i] = maxJz14(dp[i], maxJz14(i*(i-j), i*dp[i-j]))
		}
	}
	return dp[n]
}

func maxJz14(a, b int) int {
	if a < b {
		return b
	}
	return a
}
