package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)

	data1, _, _ := read.ReadLine()
	data2, _, _ := read.ReadLine()

	s := string(data1)
	t := string(data2)

	// 大写转小写
	var b strings.Builder
	var b1 strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b.WriteByte(c)
	}
	for i := 0; i < len(t); i++ {
		c := t[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b1.WriteByte(c)
	}

	ss := b.String()
	target := b1.String()

	m := make(map[string]int)
	for _, v := range ss {
		m[string(v)]++
	}
	//fmt.Printf("%+v\n", m)
	fmt.Printf("%d\n", m[target])
}
