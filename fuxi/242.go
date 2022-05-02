package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hash := make([]byte, 26, 26)

	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		hash[t[i]-'a']--
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {

	arr := make([]int, 0, 20)
	fmt.Println(arr)

	fmt.Println(isAnagram("aacc", "ccac"))
}
