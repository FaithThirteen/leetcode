package main

/**
	1.快指针先进入环
	2.慢指针进入环入口，快指针此时在环中某个位置(也可能相遇)
	3.假设此时快指针和慢指针相差x步，若在2.相遇，则x=0,
	4.设环周长为n，快指针追赶慢指针，则需要追赶n-x步，快指针每次追赶1个单位，设慢指针速度为1/s,慢指针为2/s,那么追赶需要(n-x)s
	5.在n-x秒内，慢指针走了n-x单位，因为x>=0,则慢指针小于等于n，即不走满一圈即相遇
 */

func EntryNodeOfLoop(pHead *ListNode) *ListNode{

	fast,slow := pHead,pHead

	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow{
			pre := pHead
			for pre != slow{
				pre = pre.Next
				slow = slow.Next
			}
			return pre
		}
	}

	return nil
}