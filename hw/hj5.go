package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	read := bufio.NewReader(os.Stdin)
	data,_,_ := read.ReadLine()

	v, _ := strconv.ParseInt(string(data), 0, 64)
	fmt.Println(v)

}