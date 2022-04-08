package main

import "fmt"

// dp[i] = dp[i-1] + dp[i-2],并且前两个台阶的值已经清楚，i=2，只是初始化，其实是从第三个台阶开始
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 2; i < n; i++ {
		c = a + b
		a, b = b, c
		fmt.Println("var:", c)
	}
	return c
}

func climbStairs1(n int) int {
	if n==1{
		return 1
	}
	dp:=make([]int,n+1)
	dp[1]=1
	dp[2]=2
	for i:=3;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(5))
}
