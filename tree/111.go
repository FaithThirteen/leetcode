package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	var f func(node *TreeNode,l int)int
	var length int
	f = func(node *TreeNode, l int) int {
		if node == nil {
			return l
		}
		l++
		// 有陷阱，当左侧节点为空时，如果右侧有节点，需要遍历右侧节点
		if node.Left == nil && node.Right != nil {
			return f(node.Right,l)
		}
		if node.Right == nil && node.Left != nil {
			return f(node.Left,l)
		}

		return min(f(node.Left,l),f(node.Right,l))
	}

	return f(root,length)
}

func min(a, b int) int {
	if a < b {
		return a;
	}
	return b;
}
// 递归
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0;
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right);
	}
	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left);
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1;
}

// 迭代

func minDepth2(root *TreeNode) int {
	dep := 0;
	queue := make([]*TreeNode, 0);
	if root != nil {
		queue = append(queue, root);
	}
	for l := len(queue); l > 0; {
		dep++;
		for ; l > 0; l-- {
			node := queue[0];
			if node.Left == nil && node.Right == nil {
				return dep;
			}
			if node.Left != nil {
				queue = append(queue, node.Left);
			}
			if node.Right != nil {
				queue = append(queue, node.Right);
			}
			queue = queue[1:];
		}
		l = len(queue);
	}
	return dep;
}