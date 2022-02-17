package main

import "fmt"

// pwwkew

func lengthOfLongestSubstring(s string) int {

	sb := []byte(s)
	var (
		m    = make(map[byte]struct{})
		max  int
		temp int
		l, r = 0, 0
	)

	for l < r {
		if _, ok := m[sb[r]]; ok {
			max = Max(max, temp)
			temp = 1

		} else {
			temp++
			m[sb[r]] = struct{}{}
		}
		r++
	}

	return Max(max, temp)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
