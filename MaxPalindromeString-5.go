package main

import "fmt"

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。

// 输入：s = "cbbd"
// 输出："bb"

func main() {
	s := "ac"
	res := longestPalindrome(s)
	fmt.Printf("%s \n", res)
}

// 中心扩散法
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLen := 0
	res := []int{0, 0}
	for i := 0; i < length-1; i++ {
		odd := getString(s, i, i)    // 奇数
		even := getString(s, i, i+1) // 偶数
		tempArr := []int{}
		if odd[1] > even[1] {
			tempArr = odd
		} else {
			tempArr = even
		}
		if tempArr[1] > maxLen {
			maxLen = tempArr[1]
			res = tempArr
		}
	}

	return s[res[0] : res[0]+res[1]]

}

func getString(s string, left, right int) []int {
	length := len(s)

	for left >= 0 && right < length {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	return []int{left + 1, right - left -1}
}
