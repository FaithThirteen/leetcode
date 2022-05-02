package main

// 回溯
type XY struct {
	X, Y int
}

var dir = []XY{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func hasPath(matrix [][]byte, word string) bool {

	m, n := len(matrix), len(matrix[0])

	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var dfs func(i, j, k int) bool // i,j 定位每个元素，k代表字符串下表索引
	dfs = func(i, j, k int) bool {
		if matrix[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词路径存在于网格中
			return true
		}
		used[i][j] = true                     // 标记已经使用,在子路径中不可以使用
		defer func() { used[i][j] = false }() // 回溯

		for _, d := range dir {
			if newI, newJ := i+d.X, j+d.Y; (newI >= 0 && newI < m) && (newJ >= 0 && newJ < n) && used[newI][newJ] != true {
				if dfs(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range matrix {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
