package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		s := reader.Text()
		if valid(s) {
			fmt.Printf("OK")
		} else {
			fmt.Printf("NG")
		}
	}
}

func valid(s string) bool {
	// rule1
	if len(s) <= 8 {
		return false
	}
	// rule2
	var counts [4]int
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			counts[0]++
		} else if v >= 'A' && v <= 'Z' {
			counts[1]++
		} else if v >= '0' && v <= '9' {
			counts[2]++
		} else {
			counts[3]++
		}
	}
	diffType := 0
	for _, v := range counts {
		if v != 0 {
			diffType++
		}
	}

	if diffType < 3 {
		return false
	}
	// rule3
	m := make(map[string]struct{}, len(s)/2+1)
	for i := 0; i < len(s)-2; i++ {
		if _, ok := m[s[i:i+3]]; !ok {
			m[s[i:i+3]] = struct{}{}
		} else {
			return false
		}
	}

	return true
}
