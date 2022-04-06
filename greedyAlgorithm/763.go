package main

import "fmt"

func partitionLabels(s string) []int {

	var result []int
	// 记录每个字母最大下标
	lastIndex := [26]int{}
	for i, c := range s {
		lastIndex[c-'a'] = i
	}

	start, end := 0, 0
	for i, c := range s {
		// 如果遇到字母的下标大过边界值，则更新边界值
		if lastIndex[c-'a'] > end {
			end = lastIndex[c-'a']
		}
		// 如果遇到了end,说明片段到了
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}