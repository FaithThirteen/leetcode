package main

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

// Convert
// 因为要成为有序的双向链表，因为是二叉搜索树，可以采用中序遍历，就是天然的从大到小
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
		}else{ // 如果不为nil说明已经不是最左节点
			pre.Right = cur
			cur.Left = pre
		}

		pre = cur // 更新上驱节点

		dfs(cur.Right)
	}

	dfs(pRootOfTree)

	return head
}