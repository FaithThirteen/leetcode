package main

import "container/list"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder11(root *Node) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	var tmp []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i:= 0;i< length;i++{
			node := queue.Remove(queue.Front()).(*Node)
			cLength := len(node.Children)
			for j:=0;j<cLength;j++{
				queue.PushBack(node.Children[j])
			}

			tmp = append(tmp,node.Val)
		}
		res = append(res,tmp)
		tmp = []int{}
	}
	return res
}
