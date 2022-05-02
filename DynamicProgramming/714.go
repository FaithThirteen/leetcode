package main

import "fmt"

func maxProfit(prices []int, fee int) int {

	buy := prices[0]+fee
	profit := 0
	for i:=1;i<len(prices);i++{
		if prices[i] > buy{ // 卖出
			profit += prices[i] - buy
			buy = prices[i] // 没有加上手续费是因为卖出了之后可以有一次免手续费的机会
		}else if prices[i] + fee < buy {
			buy = prices[i] + fee
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{1,3,7,5,10,3},3))
}