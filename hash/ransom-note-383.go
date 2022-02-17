package main

import "fmt"

func main() {
	ransomNote := "bg"
	magazine := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"

	fmt.Println(canConstruct(ransomNote, magazine))
}

// map解法
func canConstruct(ransomNote string, magazine string) bool {

	m := make(map[int32]int)

	for _, v := range ransomNote {
		if n, ok := m[v]; ok {
			m[v] = n + 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range magazine {
		if n, ok := m[v]; ok {
			m[v] = n - 1
		}
	}

	for _, n := range m {
		if n > 0 {
			return false
		}
	}

	return true
}
