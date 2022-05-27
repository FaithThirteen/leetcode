package main

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here

	if pRoot1 == nil || pRoot2 == nil {
		return false
	}

	// 找到相同的val
	var f func(*TreeNode) bool
	f = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if root.Val == pRoot2.Val {
			if isSame(root, pRoot2) {
				return true
			}
		}

		return f(root.Left) || f(root.Right)
	}
	return f(pRoot1)
}

func isSame(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return true
	}

	if root1 == nil && root2 == nil {
		return true
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}
