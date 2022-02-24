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
第一步：如果数组大小为零的话，说明是空节点了。

第二步：如果不为空，那么取后序数组最后一个元素作为节点元素。

第三步：找到后序数组最后一个元素在中序数组的位置，作为切割点

第四步：切割中序数组，切成中序左数组和中序右数组 （顺序别搞反了，一定是先切中序数组）

第五步：切割后序数组，切成后序左数组和后序右数组

第六步：递归处理左区间和右区间
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]

	left := firstIndex(inorder, rootVal)

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:left], postorder[:left]),
		Right: buildTree(inorder[left:], postorder[left:len(postorder)-1]),
	}

}

func firstIndex(arr []int, tar int) int {
	for i, n := range arr {
		if n == tar {
			return i
		}
	}
	return -1
}
