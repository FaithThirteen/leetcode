package main

import "fmt"

//此时无非就是要找到两个点，买入日期，和卖出日期。
//
//买入日期：其实很好想，遇到更低点就记录一下。
//卖出日期：这个就不好算了，但也没有必要算出准确的卖出日期，只要当前价格大于（最低价格+手续费），就可以收获利润，至于准确的卖出日期，就是连续收获利润区间里的最后一天（并不需要计算是具体哪一天）。
//所以我们在做收获利润操作的时候其实有三种情况：
//
//情况一：收获利润的这一天并不是收获利润区间里的最后一天（不是真正的卖出，相当于持有股票），所以后面要继续收获利润。
//情况二：前一天是收获利润区间里的最后一天（相当于真正的卖出了），今天要重新记录最小价格了。
//情况三：不作操作，保持原有状态（买入，卖出，不买不卖
func maxProfit2(prices []int, fee int) int {

	var profit int
	var minPrice = prices[0]
	for i := 1; i < len(prices); i++ {

		// 买入
		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		// 不买不卖
		if prices[i] > minPrice && prices[i] < minPrice+fee {
			continue
		}

		// 卖出
		if prices[i] > minPrice+fee {
			profit += prices[i] - minPrice - fee
			minPrice = prices[i] - fee // 当我们卖出一只股票，我们可以获得一次免手续费的权力，因为下一天股票可能还在涨，prices[i] - fee是为了减去手续费
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit2([]int{1, 3, 2, 8, 4, 9}, 2))
}

// 官方答案
func maxProfit2_(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i]+fee
		} else if prices[i] > buy {
			profit += prices[i]-buy
			buy = prices[i]
		}
	}
	return profit
}