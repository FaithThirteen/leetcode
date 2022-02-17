package main

import (
	"fmt"
)

/**

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

输入：x = 123
输出：321

输入：x = -123
输出：-321

输入：x = 120
输出：21

输入：x = 0
输出：0
*/

func main() {
	res := reverse(-123)
	fmt.Printf("result === >%d", res)
	//fmt.Println(10 ^ 3)
}

func reverse(x int) int {
	l := pow(-2, 31)
	r := pow(2, 31) - 1
	if x < l || x > r || x == 0 {
		return 0
	}
	s := 0
	arr := []int{}
	for {
		i := x % 10
		x = x / 10
		arr = append(arr, i)
		s++
		if x == 0 {
			break
		}
	}
	for _, v := range arr {
		s--
		x = x + v*pow(10, s)
	}
	if x < l || x > r || x == 0 {
		return 0
	}

	return x
}
func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
