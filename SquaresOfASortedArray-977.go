package main

import "fmt"

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

 

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
 

提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{-7,-3,2,3,11}
	fmt.Printf("%v", sortedSquares(nums))

}

func sortedSquares(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	i,j,k := 0,length-1,length-1

	for i < j || i == j {
		i2 := nums[i]*nums[i]
		j2 := nums[j]*nums[j]
		if i2 < j2 {
			result[k] = j2
			k--
			j--
		}else{
			result[k] = i2
			k--
			i++
		}
	}

	return result
}
