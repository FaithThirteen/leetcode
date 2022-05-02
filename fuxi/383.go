package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {

	m := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		count, ok := m[ransomNote[i]]
		if !ok {
			return false
		}
		if count <= 0 {
			return false
		}
		m[ransomNote[i]]--
	}
	return true
}

func main() {

	fmt.Println(canConstruct("aa","aab"))
}