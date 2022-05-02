package main



/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param threshold int整型
 * @param rows int整型
 * @param cols int整型
 * @return int整型
 */

/**
一个点只要左边或者上方可到达，并且本身满足数位和条件，那么就是可达的
*/
func movingCount(threshold int, rows int, cols int) int {
	// write code here
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	vis[0][0] = true
	var count int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if getSum(i, j) > threshold { // 如果大于k,跳过
				continue
			}
			// 如果上方可达
			if i >= 1 && vis[i-1][j] == true {
				vis[i][j] = true
			}
			// 如果上方可达
			if j >= 1 && vis[i][j-1] == true {
				vis[i][j] = true
			}
			if vis[i][j]  { // 当前格子可达
				count++
			}
		}
	}
	return count
}

func getSum(i, j int) int {
	var sum int
	for i != 0 {
		sum += i % 10
		i /= 10
	}
	for j != 0 {
		sum += j % 10
		j /= 10
	}
	return sum
}
