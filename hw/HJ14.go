package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var count int
	fmt.Scan(&count)

	read := bufio.NewReader(os.Stdin)

	var str []string
	for i := 0; i < count; i++ {
		data, _, _ := read.ReadLine()
		str = append(str, string(data))
	}

	sort.Strings(str)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%s\n", str[i])
	}

}
