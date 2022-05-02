package main

import (
	"fmt"
	"sort"
)

func main(){
	var count int
	fmt.Scan(&count)

	m := make(map[int]struct{})

	var nums []int
	for i:=0;i<count;i++{
		var n int
		fmt.Scan(&n)
		_,ok := m[n]
		if ok{
			continue
		}
		m[n] = struct{}{}
		nums = append(nums,n)
	}

	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		fmt.Printf("%d\n",nums[i])
	}
}