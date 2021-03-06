package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 定义一个dp数组
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	// 初始化, 如果是障碍物, 后面的就都是0, 不用循环了
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}
	// dp数组推导过程
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果obstacleGrid[i][j]这个点是障碍物, 那么dp[i][j]保持为0
			if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

//func main() {
//	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}))
//	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
//
//
//
//}

func consumer(cname string, ch chan int) {

	//可以循环 for i := range ch 来不断从 channel 接收值，直到它被关闭。
	for i := range ch {
		fmt.Println("consumer-----------", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
}


func main() {
	fmt.Println("i is :",Anonymous())
}

// 匿名函数
func Anonymous() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2 value is ", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1 in value is ", i)
	}()

	return i
}