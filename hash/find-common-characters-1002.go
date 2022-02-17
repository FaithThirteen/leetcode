package main

import "fmt"

func main() {
	words := []string{"bella", "label", "roller"}

	fmt.Println(commonChars(words))
}

func commonChars(words []string) []string {

	arrs := [][]int{}
	res := []string{}
	// 统计每个字符基础出现的次数
	for i := 0; i < len(words); i++ {
		arr := [26]int{}
		for j := 0; j < len(words[i]); j++ {
			arr[words[i][j]-97]++
		}

		arrs = append(arrs, arr[:])
	}

	for j := 0; j < len(arrs[0]); j++ {
		pre := arrs[0][j]
		for i := 0; i < len(arrs); i++ {
			pre = min(pre, arrs[i][j])
		}
		s := string(j + 97)
		for i := 0; i < pre; i++ {
			res = append(res, s)
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
