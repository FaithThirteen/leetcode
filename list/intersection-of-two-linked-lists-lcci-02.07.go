package main



func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]struct{})

	for headA != nil || headB != nil {
		if _,ok := m[headA];ok {
			return headA
		}
		if headA != nil {
			m[headA] = struct{}{}
			headA = headA.Next
		}


		if _,ok := m[headB];ok {
			return headB
		}
		if headB != nil {
			m[headB] = struct{}{}
			headB = headB.Next
		}
	}

	return nil
}