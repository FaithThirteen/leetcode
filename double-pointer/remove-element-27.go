package main

func removeElement(nums []int, val int) int {

	length := len(nums)

	left, right := 0, length

	for left < right {
		if nums[left] == val {
			for left < right {
				// 找到一个不等于的值进行交换
				if nums[right] == val {
					length --
					right--
					continue
				}
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
				break
			}
		}


	}

	return 0
}
