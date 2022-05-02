package main

import "fmt"


// 中心扩散法
func longestPalindrome1(s string) string {

	var maxStr string
	for i := 0; i < len(s); i++ {
		s1,s2 := getMaxStr(s,i,i),getMaxStr(s,i,i+1)
		if len(s1) > len(maxStr){
			maxStr = s1
		}
		if len(s2) > len(maxStr){
			maxStr = s2
		}
	}
	return maxStr
}

func getMaxStr(s string, i, j int) string {
	l, r := 0, len(s)

	for i >= l && j < r {
		if s[i] == s[j]{
			i--
			j++
		}else{
			break
		}
	}
	return s[i+1:j]
}

func main() {
	fmt.Println(longestPalindrome1("cbbd"))
}