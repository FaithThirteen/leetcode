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
一看到累加树，相信很多小伙伴都会疑惑：如何累加？遇到一个节点，然后在遍历其他节点累加？怎么一想这么麻烦呢。

然后再发现这是一棵二叉搜索树，二叉搜索树啊，这是有序的啊。

那么有序的元素如果求累加呢？

其实这就是一棵树，大家可能看起来有点别扭，换一个角度来看，这就是一个有序数组[2, 5, 13]，求从后到前的累加数组，也就是[20, 18, 13]，是不是感觉这就简单了。

为什么变成数组就是感觉简单了呢？

因为数组大家都知道怎么遍历啊，从后向前，挨个累加就完事了，这换成了二叉搜索树，看起来就别扭了一些是不是。

那么知道如何遍历这个二叉树，也就迎刃而解了，从树中可以看出累加的顺序是右中左，所以我们需要反中序遍历这个二叉树，然后顺序累加就可以了。
 */
func convertBST(root *TreeNode) *TreeNode {
	var f func(r *TreeNode)
	var pre = 0
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		f(r.Right)
		r.Val = r.Val+pre
		pre = r.Val
		f(r.Left)
	}
	f(root)
	return root
}