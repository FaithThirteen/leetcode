package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {

	var res int
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		// 递归终止条件
		if r == nil {
			return
		}
		// 找到符合的节点，加上值
		if r.Left != nil && r.Left.Left == nil && r.Left.Right == nil {
			res += r.Left.Val
		}
		f(r.Left)
		f(r.Right)
	}
	f(root)
	return res
}
