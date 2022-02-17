package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder1(root *TreeNode) [][]int {
	res:=[][]int{}
	if root == nil {
		return [][]int{}
	}

	tmpArr := []int{}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		for i:= 0;i<length;i++{
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr,node.Val)
		}
		res = append(res,tmpArr)
		tmpArr = []int{}
	}
	return res
}