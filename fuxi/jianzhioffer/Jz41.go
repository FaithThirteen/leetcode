package main

import (
	"sort"
)

var nums []int

func Insert(num int) {
	nums = append(nums, num)
}

func GetMedian() float64 {
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	length := len(nums)
	mid := float64(nums[length/2])
	if length%2 == 0 {
		mid = (mid + float64(nums[length/2-1])) / 2.0
	}
	//fmt.Println(mid)
	return mid
}

func main() {
	Insert(5)
	GetMedian()

	Insert(2)
	GetMedian()

	Insert(3)
	GetMedian()

	Insert(4)
	GetMedian()

	Insert(1)
	GetMedian()

	Insert(6)
	GetMedian()

	Insert(7)
	GetMedian()
}
