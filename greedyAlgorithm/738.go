package main

import (
	"fmt"
	"strconv"
)
//题目要求小于等于N的最大单调递增的整数，那么拿一个两位的数字来举例。
//
//例如：98，一旦出现strNum[i - 1] > strNum[i]的情况（非单调递增），首先想让strNum[i - 1]--，然后strNum[i]给为9，这样这个整数就是89，即小于98的最大的单调递增整数。
//
//这一点如果想清楚了，这道题就好办了。
//
//局部最优：遇到strNum[i - 1] > strNum[i]的情况，让strNum[i - 1]--，然后strNum[i]给为9，可以保证这两位变成最大单调递增整数。
//
//全局最优：得到小于等于N的最大单调递增的整数。
//
//但这里局部最优推出全局最优，还需要其他条件，即遍历顺序，和标记从哪一位开始统一改成9。
//
//此时是从前向后遍历还是从后向前遍历呢？
//
//从前向后遍历的话，遇到strNum[i - 1] > strNum[i]的情况，让strNum[i - 1]减一，但此时如果strNum[i - 1]减一了，可能又小于strNum[i - 2]。
//
//这么说有点抽象，举个例子，数字：332，从前向后遍历的话，那么就把变成了329，此时2又小于了第一位的3了，真正的结果应该是299。
//
//所以从前后向遍历会改变已经遍历过的结果！
//
//那么从后向前遍历，就可以重复利用上次比较得出的结果了，从后向前遍历332的数值变化为：332 -> 329 -> 299

func monotoneIncreasingDigits(n int) int {
	num := strconv.Itoa(n) // 将数字转为字符串，方便使用下标
	number := []byte(num)  //将字符串转为byte数组，方便更改
	l := len(number)
	for i := l - 1; i > 0; i-- {
		if number[i-1] > number[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			number[i-1] -= 1
			for j := i; j < l; j++ { // 后面的全部置为9
				number[j] = '9'
			}
		}
	}

	result, _ := strconv.Atoi(string(number))
	return result
}

func main() {
	fmt.Println(monotoneIncreasingDigits(1243))
}
