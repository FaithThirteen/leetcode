package main

import "container/list"

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		tempVal := 0
		for i:= 0;i< length;i++{
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempVal += node.Val
		}
		res = append(res,float64(tempVal)/float64(length))
	}
	return res
}