package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	var f func(r *TreeNode) int
	var result int
	f = func(r *TreeNode) int {
		// 空节点算有覆盖，则叶子节点不需要放置
		if r == nil {
			return 2
		}
		left := f(r.Left)
		right := f(r.Right)
		// 情况1：左右节点都有覆盖
		if left == 2 && right == 2 {
			return 0
		}
		// 情况2：左右节点至少有一个无覆盖的情况，需要放置摄像头
		if left == 0 || right == 0 {
			result++
			return 1
		}
		// 情况3：左右节点至少有一个有摄像头，此时已经被覆盖
		if left == 1 || right == 1 {
			return 2
		}
		// 不会走到这里
		return -1
	}
	// 情况4：头结点没有覆盖
	if f(root) == 0 {
		result++
	}
	return result
}
