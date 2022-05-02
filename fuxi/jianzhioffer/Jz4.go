package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find(target int, array [][]int) bool {

	// write code here
	// 确定列
	length := len(array)
	lenght1 := len(array[0])
	if length <= 0  || lenght1 <=0{
		return false
	}
	for i := 0; i < length; i++ {
		if target >= array[i][0] && target <= array[i][len(array[0])-1] {
			fmt.Println(array[i])
			find := BinarySearch(array[i], target)
			if find != -1 {
				return true
			}
		}
	}
	return false
}

func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(Find(7, [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}))
}
