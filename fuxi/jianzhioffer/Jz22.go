package main



// 	链表中倒数最后k个结点
func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
	// write code here
	fast,slow := pHead,pHead

	for i:=0;i<k;i++{
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}


