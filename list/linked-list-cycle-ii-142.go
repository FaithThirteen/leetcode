package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func initList(length int) *ListNode {
	if length < 1 {
		return nil
	}
	list := &ListNode{}
	head := list
	for i := 1; i <= length; i++ {
		list.Val = i
		if i != length {
			list.Next = &ListNode{}
		}
		fmt.Printf("val:%d \n", list.Val)
		list = list.Next
	}
	fmt.Println("-----------------------------")
	return head
}

func fmtList(head *ListNode) {

	for head != nil {
		fmt.Println("val:", head.Val)
		head = head.Next
	}
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
