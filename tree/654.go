package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1{
		return nil
	}
	rootVal,index := findMax(nums)
	return &TreeNode{
		Val:   rootVal,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}

func findMax(nums []int) (int,int) {
	max,index := -1,0
	for i:=0;i<len(nums);i++{
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return max,index
}
