package main

import "fmt"

func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	var res int
	for i := 1; i < len(array); i++ {
		array[i] += max(array[i-1],0) // 如果dp[i] >0 ,相加，小于0 取原先值
		res = max(res,array[i])
	}
	return res
}

func max(a,b int)int{
	if a< b{
		return b
	}
	return a
}


func main() {
	fmt.Println(FindGreatestSumOfSubArray([]int{1,-2,3,10,-4,7,2,-5}))
}