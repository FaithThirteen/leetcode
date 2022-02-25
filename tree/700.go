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
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node:=queue.Remove(queue.Front()).(*TreeNode)
			if node.Val == val{
				return node
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return nil
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val{
		return root
	}

	r1 := searchBST(root.Left,val)
	if r1 != nil {
		return r1
	}

	r2 := searchBST(root.Right,val)
	if r2 != nil {
		return r2
	}
	return nil
}
