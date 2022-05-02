package main

import(
	"fmt"
)

func main(){

	var s string
	fmt.Scan(&s)


	b := []byte(s)
	i,j:= 0,len(b)-1
	for i < j{
		b[i],b[j] = b[j],b[i]
		i++
		j--
	}
	fmt.Printf("%s\n",string(b))
}