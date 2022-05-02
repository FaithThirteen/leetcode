package main

import "fmt"

func isValid(s string) bool {

	n := len(s)
	if n%2 == 1 {
		return false
	}

	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		if m[s[i]] > 0 { // 遇到右括号弹出栈
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] { // 没有遇到匹配的括号
				return false
			}
			stack = stack[:len(stack)-1] // 弹出
		}else{ // 遇到左括号加入
			stack = append(stack,s[i])
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("{[]}"))
}