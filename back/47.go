package main

import (
	"fmt"
	"sort"
)

/**
这里开始思路我是和46一样的，唯一不同的是去重，观察了一下，同层级的不能重复选择
*/

var result47 [][]int

func permuteUnique(nums []int) [][]int {
	result47 = [][]int{}
	sort.Ints(nums)
	backtracking47(len(nums), len(nums), nums, []int{}, map[int]bool{})
	return result47
}

func backtracking47(l, n int, nums, path []int, used map[int]bool) {
	if len(path) == l {
		temp := make([]int, len(path))
		copy(temp, path)
		result47 = append(result47, temp)
		return
	}

	for i := 0; i < n; i++ {

		if i > 0 && nums[i] == nums[i-1] && used[nums[i]] == true {
			continue
		}

		path = append(path, nums[i])
		temp := make([]int, len(nums))
		copy(temp, nums)
		temp = append(temp[:i], temp[i+1:]...)

		used[nums[i]] = false
		backtracking47(l, len(temp), temp, path, used)
		used[nums[i]] = true
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1,2,3}))
}


// 答案


//var res [][]int
//func permute(nums []int) [][]int {
//	res = [][]int{}
//	backTrack(nums,len(nums),[]int{})
//	return res
//}
//func backTrack(nums []int,numsLen int,path []int)  {
//	if len(nums)==0{
//		p:=make([]int,len(path))
//		copy(p,path)
//		res = append(res,p)
//	}
//	used := [21]int{}//跟前一题唯一的区别，同一层不使用重复的数。关于used的思想carl在递增子序列那一题中提到过
//	for i:=0;i<numsLen;i++{
//		if used[nums[i]+10]==1{
//			continue
//		}
//		cur:=nums[i]
//		path = append(path,cur)
//		used[nums[i]+10]=1
//		nums = append(nums[:i],nums[i+1:]...)
//		backTrack(nums,len(nums),path)
//		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
//		path = path[:len(path)-1]
//
//	}
//
//}