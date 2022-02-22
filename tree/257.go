package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0257.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%89%80%E6%9C%89%E8%B7%AF%E5%BE%84.md
func binaryTreePaths(root *TreeNode) []string {
	// 1.确定终止条件为遇到叶子节点，为左右节点为空
	// 2.后续逻辑确保不会出现cur节点为空
	var res []string
	var travel func(r *TreeNode, s string)
	travel = func(r *TreeNode, s string) {
		if r.Left == nil && r.Right == nil {
			s = s + strconv.Itoa(r.Val)
			res = append(res, s)
		}

		s = s + strconv.Itoa(r.Val) + "->"
		if r.Left != nil {
			travel(r.Left, s)
		}
		if r.Right != nil {
			travel(r.Right, s)
		}
	}

	travel(root, "")
	return res
}
