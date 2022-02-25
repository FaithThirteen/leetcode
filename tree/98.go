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
func isValidBST(root *TreeNode) bool {

	var arr []int
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		f(r.Left)
		arr = append(arr,r.Val)
		f(r.Right)
	}
	f(root)
	return Check(arr)
}

func Check(arr []int)bool{
	for i:=1;i<len(arr);i++{
		if arr[i-1] >= arr[i]{
			return false
		}
	}
	return true
}

func isValidBST1(root *TreeNode) bool {
	// 二叉搜索树也可以是空树
	if root == nil {
		return true
	}
	// 由题目中的数据限制可以得出min和max
	return check(root,math.MinInt64,math.MaxInt64)
}

func check(node *TreeNode,min,max int64) bool {
	if node == nil {
		return true
	}

	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}
	// 分别对左子树和右子树递归判断，如果左子树和右子树都符合则返回true
	return check(node.Right,int64(node.Val),max) && check(node.Left,min,int64(node.Val))
}

// 中序遍历解法
func isValidBST2(root *TreeNode) bool {
	// 保存上一个指针
	var prev *TreeNode
	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := travel(node.Left)
		// 当前值小于等于前一个节点的值，返回false
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}
	return travel(root)
}