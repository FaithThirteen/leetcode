package main


/**
 * ps: 今天比较有感觉
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	res := &TreeNode{}
	var traversal func(r *TreeNode,tar *TreeNode)
	traversal = func(r *TreeNode,tar *TreeNode){
		if r == nil{
			return
		}
		tar.Val = r.Val
		if r.Right != nil {
			tar.Left = &TreeNode{}
		}
		if r.Left != nil {
			tar.Right = &TreeNode{}
		}
		traversal(r.Right,tar.Left)
		traversal(r.Left,tar.Right)
	}

	traversal(root,res)
	return res
}

func invertTree1(root *TreeNode) *TreeNode {
	if root ==nil{
		return nil
	}
	temp:=root.Left
	root.Left=root.Right
	root.Right=temp

	invertTree1(root.Left)
	invertTree1(root.Right)

	return root
}