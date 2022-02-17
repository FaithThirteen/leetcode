package main


func reverseList(head *ListNode) *ListNode {

	var temp *ListNode
	var newHead *ListNode

	for head != nil {
		temp = newHead
		newNode := &ListNode{Val: head.Val, Next: temp}
		newHead = newNode
		head = head.Next
	}

	return newHead
}

func main() {

	list := &ListNode{}
	head := list
	for i := 0; i < 8; i++ {
		list.Val = i
		if i != 7 {
			list.Next = &ListNode{}
		}
		list = list.Next
	}

	revList := reverseList(head)
	fmtList(revList)

}
