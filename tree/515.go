package main

import (
	"container/list"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		max := int(math.Pow(-2,31))
		for i := 0;i < length;i++{
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Val > max{
				max = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res,max)
		max = 0
	}


	return res
}