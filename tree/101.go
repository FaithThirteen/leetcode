package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func defs(left *TreeNode, right *TreeNode) bool {
	// 两棵树都为空，对称，无需判断值，返回true
	if left == nil && right == nil {
		return true
	}
	// 其中有一个为空，说明不对称
	if left == nil || right == nil {
		return false
	}
	// 此时两棵树的节点都有值，判断是否值相等，相等才对称
	if left.Val != right.Val {
		return false
	}
	// 左树的左侧等于右树的右侧，左树的右侧等于右树的左侧，才可以算对称
	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
}
func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}