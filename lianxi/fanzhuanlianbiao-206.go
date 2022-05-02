package main


import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // 记录下next，方便将curr移动
		// 为prev新增一个头节点
		curr.Next = prev
		prev = curr

		curr = next
	}
	return prev
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	newHead := reverseList(head)

	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}

}
