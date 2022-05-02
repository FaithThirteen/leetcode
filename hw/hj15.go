package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)

	num2 := strconv.FormatInt(int64(num),2)
	count := 0
	for len(num2) > 0 {
		if num2[len(num2)-1] == '1' {
			count++
		}
		num2 = num2[:len(num2)-1]
	}

	fmt.Println(count)

}