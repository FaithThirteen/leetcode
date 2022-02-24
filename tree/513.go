package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {

	queue := list.New()
	queue.PushBack(root)
	res := root.Val
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				res = node.Val
			}
			if node.Left!= nil {
				queue.PushBack(node.Left)
			}
			if node.Right!= nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
