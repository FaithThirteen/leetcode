package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	s := string(data)

	fmt.Println("s:", s)

	if len(s) == 0 {
		fmt.Printf("%d\n", 0)
		return
	}
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length == 0 {
			continue
		} else if s[i] == ' ' && length != 0 {
			break
		}
		// 不为空
		length++
	}
	fmt.Printf("%d\n", length)
}
