package main

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

func swapPairs(head *ListNode) *ListNode {

	var result = &ListNode{
		Next: head,
	}

	// 虚拟头节点
	pre := result
	for head != nil && head.Next != nil {
		// 1 2 3 4
		// 此时虚拟头指向2，next保存2后续正常节点
		pre.Next = head.Next
		next := head.Next.Next
		// head.Next.Next即为2的下一个为，此时指向1
		// head.Next即为1的下一个为，即刚保存有的正常节点2的后续 next
		head.Next.Next = head
		head.Next = next

		// 虚拟头节点移动，此时pre.Next = 3，head变为3
		pre = head
		head = next
	}

	return result.Next
}


func main() {

	head := initList(4)
	result := swapPairs(head)

	fmtList(result)

}
