package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 判断左右子树是否平衡，如不平衡提前结束，如两树都平衡则计算两树是否绝对值大于1
	if !isBalanced(root.Left) || !isBalanced(root.Right){
		return false
	}
	if abs(maxDepth1(root.Left)-maxDepth1(root.Right)) > 1{
		return false
	}
	return true
}

func maxDepth1(root *TreeNode) int {
	var f func(node *TreeNode,l int)int
	var length int
	f = func(node *TreeNode, l int) int {
		if node == nil {
			return l
		}
		l++
		return max(f(node.Left,l),f(node.Right,l))
	}

	return f(root,length)
}

func abs(a int)int{
	if a < 0{
		return -a
	}
	return a
}

