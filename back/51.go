package main

import (
	"fmt"
	"strings"
)

var result51 [][]string
var checkerboard [][]string

func solveNQueens(n int) [][]string {
	result51 = [][]string{}
	checkerboard = make([][]string, n)
	for i := 0; i < n; i++ {
		checkerboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			checkerboard[i][j] = "."
		}
	}

	backtracking51(0, n)

	return result51
}

func backtracking51(row, n int) {
	if row == n {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(checkerboard[i], "")
		}
		result51 = append(result51, temp)
		return
	}

	for i := 0; i < n; i++ {
		if !Validata(row, i, n) {
			continue
		}
		checkerboard[row][i] = "Q"
		backtracking51(row+1, n)
		checkerboard[row][i] = "."
	}

}

func Validata(row, i, n int) bool {
	// 验证90°上测无皇后
	for j := row - 1; j >= 0; j-- {
		if checkerboard[j][i] == "Q" {
			return false
		}
	}

	// 验证180°无皇后
	for j := 0; j < n; j++ {
		if j == i {
			continue
		}
		if checkerboard[row][j] == "Q" {
			return false
		}
	}
	// 验证45°无皇后
	r, j := row-1, i-1
	for r >= 0 && j >= 0 {
		if checkerboard[r][j] == "Q" {
			return false
		}
		r--
		j--
	}

	// 验证135°无皇后
	r, j = row-1, i+1
	for r >= 0 && j < n {
		if checkerboard[r][j] == "Q" {
			return false
		}
		r--
		j++
	}

	return true
}

func main() {

	fmt.Println(solveNQueens(4))
}
