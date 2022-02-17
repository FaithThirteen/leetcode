package main

import "fmt"

func main() {

	s := "abcdefg"
	n := 2

	fmt.Printf("%s \n", reverseLeftWords(s, n))


	fmt.Printf("%s \n", reverseLeftWords2(s, n))


}

func reverseLeftWords(s string, n int) string {
	ss := []byte(s)
	return string(ss[n:]) + string(ss[:n])
}

// 解法2，不使用额外的空间
func reverseLeftWords2(s string, n int) string {
	if n < 1{
		return s
	}

	ss := []byte(s)
	// 反转整个字符串，再反转0-n与n-len(s)
	reverseArea(ss, 0, len(ss)-1)
	reverseArea(ss, 0, len(ss)-1-n)
	reverseArea(ss, len(ss)-n, len(ss)-1)


	return string(ss)
}

func reverseArea(s []byte, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
