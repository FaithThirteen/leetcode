package main

import "fmt"

func candy(ratings []int) int {

	if len(ratings) < 2 {
		return 1
	}

	candyBag := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candyBag[i] = 1
	}
	// 右边分数大的先获取
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyBag[i] = candyBag[i-1] + 1
		}
	}

	// 左边分数大的先获取
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			candyBag[i-1] = findMax(candyBag[i-1], candyBag[i]+1) // 只有取最大的才能既保持对左边糖果多，也比右边candyBag[i + 1]的糖果多。
		}
	}

	var count int
	for _, c := range candyBag {
		count += c
	}

	return count
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	fmt.Println(candy([]int{1, 0, 2}))
}
