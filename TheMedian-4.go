package main

import "fmt"

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2

func main() {

	n1 := []int{2}
	n2 := []int{}

	result := findMedianSortedArrays(n1, n2)
	fmt.Println("中位数=====>", result)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	i := 0
	j := 0
	array := []int{}

	for i < m && j < n {
		if nums1[i] > nums2[j] {
			array = append(array, nums2[j])
			j++
		} else {
			array = append(array, nums1[i])
			i++
		}
	}
	if i < m {
		array = append(array, nums1[i:]...)
	}
	if j < n {
		array = append(array, nums2[j:]...)
	}

	length := len(array)

	if length%2 == 1 {
		return float64(array[length/2])
	}

	return (float64(array[length/2-1]) + float64(array[length/2])) / 2
}
