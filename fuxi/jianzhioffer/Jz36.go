package main

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param pRootOfTree TreeNode类
 * @return TreeNode类
 */
func Convert( pRootOfTree *TreeNode ) *TreeNode {
	// write code here
	var head,pre *TreeNode

	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		dfs(cur.Left)
		if pre == nil { // 此时pre为null,这是最左节点,即新的节点
			head = cur
		}else{ //
			pre.Right = cur
			cur.Left = pre
		}

		pre = cur

		dfs(cur.Right)
	}

	dfs(pRootOfTree)

	return head
}