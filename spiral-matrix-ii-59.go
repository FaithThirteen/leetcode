package main

import "fmt"

/**
59. 螺旋矩阵 II
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：
1 2 3
8 9 4
7 6 5

输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]
*/

func main() {

	fmt.Println(generateMatrix(5))

}
func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	m := n*n
	result := make([][]int, n)
	for i:= 0;i<n;i++{
		result[i] = make([]int,n)
	}
	for num <= m {
		// 填充上边界
		for i:= left;i<= right;i++{
			result[top][i] = num
			num++
		}
		top++
		// 填充右边界
		for i:= top;i<= bottom;i++{
			result[i][bottom] = num
			num++
		}
		right--
		// 填充下边界
		for i:= right;i>=left;i--{
			result[bottom][i] = num
			num++
		}
		bottom--
		// 填充左边界
		for i:= bottom;i>=top;i--{
			result[i][left] = num
			num++
		}
		left++
	}
	return result
}
