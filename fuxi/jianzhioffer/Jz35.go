package main

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

/**
 *
 * @param pHead RandomListNode 类
 * @return RandomListNode类
 */
func Clone(head *RandomListNode) *RandomListNode {

    if head == nil {
        return nil
    }

    p1,p2 := head,head
    m := make(map[*RandomListNode]*RandomListNode)


    for p1 != nil {
    	m[p1] = &RandomListNode{Label: p1.Label}
    	p1 = p1.Next
	}

	for p2 != nil {
		m[p2].Next = m[p2.Next]
		m[p2].Random = m[p2.Random]
		p2 = p2.Next
	}

	return m[head]
}

