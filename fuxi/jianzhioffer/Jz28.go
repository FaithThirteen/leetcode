package main

func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot ==nil{
		return true
	}
	// write code here

	var f func(r1, r2 *TreeNode) bool
	f = func(r1, r2 *TreeNode) bool {
		if r1 == nil && r2 != nil {
			return false
		}
		if r1 != nil && r2 == nil {
			return false
		}
		if r1 ==nil && r2 == nil {
			return true
		}

		if r1.Val != r2.Val{
			return false
		}

		return f(r1.Left,r2.Right) && f(r1.Right,r2.Left)
	}

	return f(pRoot.Left,pRoot.Right)
}
