package main

import "fmt"

func solveSudoku(board [][]byte) {
	backtracking37(board)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if j != 8 {
				fmt.Printf("%s", string(board[i][j]))
			} else {
				fmt.Printf("%s \n", string(board[i][j]))
			}
		}
	}
}

func backtracking37(board [][]byte) bool{

	for i := 0; i < 9; i++ {        // 遍历行
		for j := 0; j < 9; j++ {    // 遍历列
			if board[i][j] != '.' { // 是否可以放置
				continue
			}

			//尝试填1-9
			for k:='1';k<='9';k++{
				if valid(i, j,byte(k), board) {  // (i, j) 这个位置放k是否合适
					board[i][j]=byte(k)
					if  backtracking37(board)==true{  // 找到合适的立马返回
						return true
					}
					board[i][j]='.' // 说明分支不行，回溯
				}
			}
			return false  // 9个数都试完了，都不行，那么就返回false
		}
	}

	return true // 遍历无返回false，说明找到了
}

func valid(row, col int,n byte, board [][]byte) bool {
	// 判断行是否存在
	for k := 0; k < 9; k++ {
		if board[row][k] == n {
			return false
		}
	}

	// 判断列是否存在
	for k := 0; k < 9; k++ {
		if board[k][col] == n {
			return false
		}
	}

	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == n {
				return false
			}
		}
	}

	return true
}

func main() {
	solveSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
}
