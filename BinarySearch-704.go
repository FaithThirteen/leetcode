package main

import "fmt"

func main() {

	arr := []int{-1,0,3,5,9,12}
	target := 2
	fmt.Println("result : ", search(arr, target))

}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	result := -1
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		result = search(nums[:mid], target)
	}

	if nums[mid] < target {
		result = search(nums[mid:], target)
		if result != -1{
			return mid + result
		}
	}

	return result
}
