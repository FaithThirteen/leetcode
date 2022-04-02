package main

import "fmt"

// 假如第0天买入，第3天卖出，那么利润为：prices[3] - prices[0]。
//
//相当于(prices[3] - prices[2]) + (prices[2] - prices[1]) + (prices[1] - prices[0])。
//
//此时就是把利润分解为每天为单位的维度，而不是从0天到第3天整体去考虑！
//
//那么根据prices可以得到每天的利润序列：(prices[i] - prices[i - 1]).....(prices[1] - prices[0])。
func maxProfit(prices []int) int {

	var result int
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			result += profit
		}
	}
	return result
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}