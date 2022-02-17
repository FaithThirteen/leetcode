package main

import (
	"fmt"
	"math"
)

/**
输入：s = "   -42"
输出：-42
解释：
第 1 步："   -42"（读入前导空格，但忽视掉）
            ^
第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
             ^
第 3 步："   -42"（读入 "42"）
               ^
解析得到整数 -42 。
由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42 。
示例 3：

输入：s = "4193 with words"
输出：4193
解释：
第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
             ^
解析得到整数 4193 。
由于 "4193" 在范围 [-231, 231 - 1] 内，最终结果为 4193 。
示例 4：

输入：s = "words and 987"
输出：0
解释：
第 1 步："words and 987"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："words and 987"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："words and 987"（由于当前字符 'w' 不是一个数字，所以读入停止）
         ^
解析得到整数 0 ，因为没有读入任何数字。
由于 0 在范围 [-231, 231 - 1] 内，最终结果为 0 。
示例 5：

输入：s = "-91283472332"
输出：-2147483648
解释：
第 1 步："-91283472332"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："-91283472332"（读入 '-' 字符，所以结果应该是负数）
          ^
第 3 步："-91283472332"（读入 "91283472332"）
                     ^
解析得到整数 -91283472332 。
由于 -91283472332 小于范围 [-231, 231 - 1] 的下界，最终结果被截断为 -231 = -2147483648 。

*/

// 状态机
// 是数字，‘ ’，'+','-'，或者其他
func main() {
	//
	str := "-5-"

	fmt.Println(myAtoi(str))

}

func myAtoi(s string) int {
	numMap := map[int32]int{43: 1, 45: -1, 48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}
	arr := []int{}
	i := 0
	start := false // 开始读取
	flag := 0
	space := false
	other := false // 遇到其他符号
	m := false     // 负数
	//end := false   // 结束
	for _, v := range s {
		status := Status(v)
		switch status {
		case 0: // 空格
			if start {
				space = true
			}
			continue
		case 1: // 正负号
			if flag != 0 || start || other {
				return 0
			}
			if v == 45 {
				m = true
			}
			flag = int(v)
			if start || other {
				return 0
			}
			continue
		case 2: // 数字
			if other { // 如果遇到其他符号且结束标志为1
				if start {
					break
				}
				return 0
			}
			if space {
				break
			}
			n, _ := numMap[v]
			arr = append(arr, n)
			i++
			start = true
		case 3:
			if !start { // 没有开启start
				return 0
			}
			other = true
			break
		}
	}

	var result float64
	l := pow2(-2, 31)
	h := pow2(2, 31) - 1
	for _, v := range arr {
		i--
		result = result + float64(v) * math.Pow10(i)
		if result > h && !m {
			return int(h)
		}
		if result < l && m {
			return int(l)
		} else if result < l && !m {
			return int(h)
		}
	}

	if m {
		a := (result * -1)
		if a < l {
			return int(l)
		}

		return int(a)
	}

	if result > h {
		return int(h)
	}

	return int(result)
}
func pow2(x float64, n int) float64 {
	ret := 1.0 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

func Status(v int32) int {
	// ' '
	if v == 32 {
		return 0
	}
	// + -
	if v == 43 || v == 45 {
		return 1
	}
	// 0-9
	if v > 47 && v < 58 {
		return 2
	}
	// 其他符号
	return 3
}
