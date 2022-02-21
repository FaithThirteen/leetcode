package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	var f func(r *TreeNode)
	var length int
	f  = func(r *TreeNode) {
		if r == nil {
			return
		}
		length ++
		f(r.Left)
		f(r.Right)
	}
	f(root)
	return length
}

//本题直接就是求有多少个节点，无脑存进数组算长度就行了。
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}