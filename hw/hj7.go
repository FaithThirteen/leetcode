package main

import (
	"fmt"
)

func main(){
	var n float64
	fmt.Scan(&n)

	if n - float64(int(n)) < 0.5{
		fmt.Printf("%D\n",float64(int(n)))
	}else{
		fmt.Printf("%f\n",float64(int(n)+1))
	}
}