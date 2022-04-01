package main

import "fmt"

func wiggleMaxLength(nums []int) int {

	var result = 1
	var branch = 0
	for i := 0; i < len(nums); i++ {
		// 决定开头
		if i > 0 && branch == 0 {
			if nums[i]-nums[i-1] > 0 {
				branch = 1
				result++
			} else if nums[i]-nums[i-1] < 0 {
				branch = 2
				result++
			}
			continue
		}
		switch branch {
		case 1: // 上一个为正，需为负
			if nums[i]-nums[i-1] < 0 {
				result++
				branch = 2
			}
		case 2: // 上一个为负，需为正
			if nums[i]-nums[i-1] > 0 {
				result++
				branch = 1
			}
		}
	}
	return result
}


func wiggleMaxLength1(nums []int) int {
	var count,preDiff,curDiff  int
	count=1
	if len(nums)<2{
		return count
	}
	for i:=0;i<len(nums)-1;i++{
		curDiff=nums[i+1]-nums[i]
		//如果有正有负则更新下标值||或者只有前一个元素为0（针对两个不等元素的序列也视作摆动序列，且摆动长度为2）
		if (curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0){
			preDiff=curDiff
			count++
		}
	}
	return count
}

func main() {

	fmt.Println(wiggleMaxLength([]int{3, 3, 3, 2, 5}))
}
