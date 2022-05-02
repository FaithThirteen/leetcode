package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			sum[a+b]++
		}
	}

	var count int
	for _, c := range nums3 {
		for _, d := range nums4 {
			if v, ok := sum[0-(c+d)]; ok {
				count += v
			}
		}
	}

	return count
}

func main() {
	fmt.Println(fourSumCount([]int{1,2},[]int{-2,-1},[]int{-1,2},[]int{0,2}))
}
