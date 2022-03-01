package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 暴力解法
func findMode(root *TreeNode) []int {
	history := make(map[int]int)
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		f(r.Left)
		if _, ok := history[r.Val]; !ok {
			history[r.Val] = 1
		} else {
			history[r.Val]++
		}
		f(r.Right)
	}
	f(root)

	var maxValue int
	var res []int
	for _, v := range history {
		if v > maxValue{
			maxValue = v
		}
	}

	for k, v := range history {
		if maxValue == v{
			res = append(res,k)
		}
	}

	return res
}

func findMode1(root *TreeNode) []int {

	var pre *TreeNode
	var f func(r *TreeNode)
	var max = 1
	var count int
	var res []int
	f = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		f(cur.Left)
		if pre != nil && pre.Val == cur.Val{
			count ++
		}else{
			count = 1
		}
		if count > max {
			res = []int{cur.Val}
			max = count
		}else if count == max{
			res = append(res,cur.Val)
		}
		pre = cur
		f(cur.Right)
	}
	f(root)
	return res
}
