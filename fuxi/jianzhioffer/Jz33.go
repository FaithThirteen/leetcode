package main

func verifyPostorder(postorder []int) bool {

	var dfs func(int, int) bool
	dfs = func(i int, j int) bool {
		if i >= j { // 节点小于等于1,无需判断
			return true
		}
		p := i // 左子树从i开始
		for postorder[p] < postorder[j] {
			p++
		}

		mid := p // 左右字数分割点
		for postorder[p] > postorder[j] {
			p++
		}
		return p == j && dfs(i, mid-1) && dfs(mid, j-1)
	}
	return dfs(0, len(postorder)-1)
}
