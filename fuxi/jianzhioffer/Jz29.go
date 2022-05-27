package main

import "fmt"

func printMatrix(matrix [][]int) []int {
	// write code here
	left, right := 0, len(matrix[0])-1
	down, up := 0, len(matrix)-1

	result := make([]int, 0, (right+1)*(up+1))
	for left <= right && down <= up {

		for i := left; i <= right; i++ {
			result = append(result, matrix[down][i])
		}
		down++
		if down > up {
			break
		}

		for i := down; i <= up; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[up][i])
		}
		up--
		if down > up {
			break
		}

		for i := up; i >= down; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}

	}
	return result
}

func main() {
	fmt.Println(printMatrix([][]int{{1}, {2}, {3}, {4}, {5}}))
}
