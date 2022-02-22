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
// 初步想法是层序遍历 + 比较两树是否相同
// https://leetcode-cn.com/problems/subtree-of-another-tree/solution/ling-yi-ge-shu-de-zi-shu-by-leetcode-solution/
// 看了题解，重新定义了简单题
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i:=0;i<length;i++{
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Val == subRoot.Val{
				if isSame(node,subRoot) {
					return true
				}
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return false
}

func isSame(p,q *TreeNode)bool{
	var queue []*TreeNode
	if p != nil || q != nil {
		queue = append(queue,p,q)
	}
	for len(queue) > 0{
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left ==nil || right ==nil || left.Val != right.Val {
			return false
		}
		queue = append(queue,left.Left,right.Left,left.Right,right.Right)
	}

	return true
}