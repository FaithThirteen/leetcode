package main

import (
	"bytes"
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)
	var b bytes.Buffer
	for {
		temp := n % 10
		b.WriteString(fmt.Sprintf("%d",temp))
		n = n /10
		if n == 0 {
			break
		}
	}
	fmt.Println(b.String())
}
