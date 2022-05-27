package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 虚拟头节点
func deleteNode(head *ListNode, val int) *ListNode {
	// write code here
	newHead := &ListNode{
		Next: head,
	}
	cur:= newHead

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return newHead.Next
}
