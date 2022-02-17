package main

import "fmt"

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isAnagram(s string, t string) bool {

	m := make(map[int32]int)

	// 放入map
	for _, v := range s {
		n, ok := m[v]
		if !ok {
			m[v] = 1
			continue
		}
		m[v] = n + 1
	}

	for _, v := range t {
		n, ok := m[v]
		if ok {
			m[v] = n - 1
		}else{
			return false
		}
	}

	for _, v := range m {
		// 有一个大于0则说明有字母出现的次数不为0
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "rat"
	t := "car"

	fmt.Println(isAnagram(s,t))
}