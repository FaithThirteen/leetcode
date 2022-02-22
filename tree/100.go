package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var queue []*TreeNode

	if p != nil || q != nil{
		queue = append(queue,p,q)
	}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val  {
			return false
		}
		queue = append(queue,left.Left,right.Left,left.Right,right.Right)
	}
	return true
}