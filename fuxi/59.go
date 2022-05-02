package main

import "fmt"

func generateMatrix(n int) [][]int {

	num, target := 1, n*n
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	left, right := 0, n-1
	top, bottom := 0, n-1

	for num <= target {

		// 填充左到右
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		top++

		// 填充右到下
		for i := top; i <= bottom; i++ {
			result[i][right] = num
			num++
		}

		right--
		// 填充右到左
		for i := right; i >= left; i-- {
			result[bottom][i] = num
			num++
		}

		bottom--
		// 填充下到上
		for i := bottom; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++

	}

	return result
}

func main() {
	fmt.Println(generateMatrix(3))
}
