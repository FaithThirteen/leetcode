package main

import "fmt"



// https://camo.githubusercontent.com/0dec5b93fd4de24b7e08293f8f52f6cebd689c2c9073a656c01f33c1e4c16890/68747470733a2f2f696d672d626c6f672e6373646e696d672e636e2f32303230313131383230323434383634322e706e67

var result131 [][]string

func partition(s string) [][]string {
	result131 = [][]string{}
	backtracking131(0,s,[]string{})
	return result131
}

func backtracking131(start int, str string, path []string) {
	if start == len(str) {
		temp := make([]string, len(path))
		copy(temp, path)
		result131 = append(result131, temp)
	}

	for i := start; i < len(str); i++ {
		tempStr := str[start:i+1]  // 对余下的字符串，一一枚举，如果取得的字符串为回文，则继续
		if !backStr(tempStr) { // 截取的不为回文,跳过
			continue
		}
		path = append(path, tempStr)
		backtracking131(i+1, str, path)
		path = path[:len(path)-1]
	}

}

func backStr(s string) bool {
	if len(s) < 1 {
		return false
	}
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}