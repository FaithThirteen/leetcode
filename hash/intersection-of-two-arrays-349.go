package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		if _, ok := m[nums1[i]]; ok {
			continue
		}
		m[nums1[i]] = 1
	}

	for i := 0; i < len(nums2); i++ {
		if n, ok := m[nums2[i]]; ok {
			if n > 0 {
				res = append(res, nums2[i])
				m[nums2[i]] = n - 1
			}
		}
	}

	return res
}
