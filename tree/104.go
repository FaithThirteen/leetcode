package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var f func(node *TreeNode, l int) int
	var length int
	f = func(node *TreeNode, l int) int {
		if node == nil {
			return l
		}
		l++
		l1 := f(node.Left, l)
		l2 := f(node.Right, l)

		return max(l1, l2)
	}

	return f(root, length)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归
func maxdepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxdepth(root.Left), maxdepth(root.Right)) + 1
}

// 遍历
func maxdepth1(root *TreeNode) int {
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}

func maxDepth2(root *TreeNode) int {
	var f func(r *TreeNode)int
	f = func(r *TreeNode)int {
		if r == nil {
			return 0
		}
		return max(f(r.Left),f(r.Right))+1
	}

	return f(root)
}
