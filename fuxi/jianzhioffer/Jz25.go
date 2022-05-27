package main

// 因为只在开头创建了preHead 与 pre 其余都是引用原有的内存地址，所以空间复杂度为o(1)

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	preHead := &ListNode{Next: pHead1}
	pre := preHead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val{
			pre.Next = pHead1
			pHead1 = pHead1.Next
		}else{
			pre.Next = pHead2
			pHead2 = pHead2.Next
		}
		pre = pre.Next
	}

	if pHead1 != nil {
		pre.Next = pHead1
	}
	if pHead2 != nil {
		pre.Next = pHead2
	}

	return preHead.Next
}
