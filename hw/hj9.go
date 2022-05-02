package main

import(
	"fmt"
)

func main(){
	var n string
	fmt.Scan(&n)


	m := make(map[uint8]struct{})
	var result string
	for i:=len(n)-1;i>=0;i--{
		if _,ok := m[n[i]];ok{
			continue
		}
		m[n[i]] = struct{}{}
		result += string(n[i])
	}
	fmt.Println(result)
}