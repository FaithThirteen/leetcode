package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {

	var f func(r *TreeNode)
	var arr []int
	f = func(r *TreeNode) {
		if r == nil {return}
		f(r.Left)
		arr = append(arr,r.Val)
		f(r.Right)
	}
	f(root)
	min := 1000000
	for i:=1;i<len(arr);i++{
		temp := arr[i] - arr[i-1]
		if  temp < min {
			min = temp
		}
	}
	return min
}

// 中序遍历的同时计算最小值
func getMinimumDifference1(root *TreeNode) int {
	// 保留前一个节点的指针
	var prev *TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val - prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}

