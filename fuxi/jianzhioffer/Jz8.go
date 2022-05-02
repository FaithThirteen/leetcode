package main


type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	} // 找到根节点
	var head *TreeLinkNode
	var target = pNode
	cur := pNode
	for cur != nil {
		head = cur
		cur = cur.Next
	}
	// 中序遍历
	var result []*TreeLinkNode
	var f func(linkNode *TreeLinkNode)
	f = func(r *TreeLinkNode) {
		if r == nil {
			return
		}
		f(r.Left)
		result = append(result, r)
		f(r.Right)
	}
	// 开始填充result
	f(head)

	for i := 1; i < len(result); i++ {
		if result[i-1] == target {
			return result[i]
		}
	}
	return nil
}
