package main


func main() {

	head := initList(5)
	newHead := removeNthFromEnd_1(head, 5)

	fmtList(newHead)

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast := head
	slow := &ListNode{
		Next: head,
	}

	count := 0
	for fast != nil {
		fast = fast.Next
		count++
	}

	length := count
	//fmt.Printf("count :%d \n", count)
	for slow.Next != nil {
		if count == n {
			slow.Next = slow.Next.Next
			if length == count { // 在删除的为第一个节点时，因为虚拟头节点的存在导致删除不了,特别加了这个判断
				return slow.Next
			}
			break
		}
		count--
		slow = slow.Next
	}

	return head
}

/**
让cur先走n步后pre和cur同时走，那么prev少走了n步，即为需要删除的节点，牛逼
*/
func removeNthFromEnd_1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
