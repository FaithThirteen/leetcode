package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
有以下五种情况：

第一种情况：没找到删除的节点，遍历到空节点直接返回了
找到删除的节点
第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果遇到删除的节点
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil && root.Right != nil {
			// 左孩子为空，右孩子不为空，返回右孩子
			return root.Right
		} else if root.Left != nil && root.Right == nil {
			// 左孩子不为空，右孩子为空，返回左孩子
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil { // 找到右孩子的最左边叶子节点，将当前节点的左节点拼接上去
				cur = cur.Left
			}
			cur.Left = root.Left
			root.Left = nil
			return root.Right
		}
	}

	if root.Val > key{
		root.Left = deleteNode(root.Left,key)
	}

	if root.Val < key{
		root.Right = deleteNode(root.Right,key)
	}

	return root
}
