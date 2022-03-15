package main

import (
	"fmt"
	"sort"
)

var result90 [][]int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result90 = [][]int{[]int{}}
	backtracking90(len(nums),0,nums,[]int{}, map[int]bool{})
	return result90
}

func backtracking90(n,start int,nums,path []int,used map[int]bool){

	for i:=start;i<n;i++{
		// 剪枝 去除重复集合
		if i >0 && nums[i] == nums[i-1] && used[i-1]==true{
			continue
		}

		path  = append(path,nums[i])
		temp := make([]int,len(path))
		copy(temp,path)
		result90 = append(result90,temp)
		used[i] = false // 同树枝可以使用
		backtracking90(n,i+1,nums,path,used)
		path = path[:len(path)-1]
		used[i]=true // 同树层不可在使用
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1,2,2}))
}