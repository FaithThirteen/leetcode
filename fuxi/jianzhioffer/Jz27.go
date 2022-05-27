package main

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return nil
	}
	left := Mirror(pRoot.Left)
	right := Mirror(pRoot.Right)
	pRoot.Left = right
	pRoot.Right = left
	return pRoot
}
