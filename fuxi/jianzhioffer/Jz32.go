package main

import "container/list"

func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here

	if root == nil {
		return nil
	}
	var result []int
	var l list.List
	l.PushFront(root)

	for l.Len() != 0 {
		length := l.Len()
		for i := 0; i < length; i++ {
			node := l.Front().Value.(*TreeNode)
			l.Remove(l.Front())
			result = append(result, node.Val)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
	}
	return result
}

func PrintFromTopToBottom1(root *TreeNode) []int {
	// write code here

	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue  = queue[1:]
			result = append(result, node.Val)
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
	}
	return result
}
