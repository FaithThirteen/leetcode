package main

import "fmt"

// 假设 n 个节点存在二叉排序树的个数是 G (n)，令 f(i) 为以 i 为根的二叉搜索树的个数
// G(n) = f(1)+f(2)+f(3)+f(4)+...+f(n)
// f(n) = G(i-1) * G(n-i)
// G(n)=G(0)∗G(n−1)+G(1)∗(n−2)+...+G(n−1)∗G(0)

func numTrees(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			//对于第i个节点，需要考虑1作为根节点直到i作为根节点的情况，所以需要累加
			//一共i个节点，对于根节点j时,左子树的节点个数为j-1，右子树的节点个数为i-j
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(5))
}