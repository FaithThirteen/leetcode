package main

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return string字符串一维数组
 */

var result []string

func Permutation(str string) []string {

	s := []byte(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	// write code here
	backtracking(len(str), len(str), map[int]bool{}, "", string(s))
	return result
}

func backtracking(end, resLen int, used map[int]bool, path, s string) {
	if len(path) == resLen {
		result = append(result, path)
		return
	}

	for i := 0; i < end; i++ {
		if i>0 && s[i]==s[i-1] && used[i-1]==false { // 树层去重
			continue
		}
		if used[i] == true{ // 如果该字母使用过
			continue
		}
		path = path + string(s[i])
		used[i] = true
		backtracking(end, resLen, used, path, s)
		used[i] = false
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(Permutation("abc"))
}
