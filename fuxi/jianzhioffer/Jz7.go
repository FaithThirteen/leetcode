package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: pre[0],
	}
	// 找出根节点
	i := 0
	for ; i < len(vin); i++ {
		if pre[0] == vin[i]{
			break
		}
	}

	root.Left = reConstructBinaryTree(pre[1:len(vin[:i])+1],vin[:i])
	root.Right = reConstructBinaryTree(pre[len(vin[:i])+1:],vin[i+1:])
	return root
}
