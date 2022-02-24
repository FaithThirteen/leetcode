package main



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <1 || len(inorder) < 1{
		return nil
	}
	rootVal := preorder[0]
	left := firstIndex(inorder,rootVal)

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree1(preorder[1:left+1], inorder[:left+1]),
		Right: buildTree1(preorder[left+1:], inorder[left+1:]),
	}
}