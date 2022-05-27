package main

func FindPath(root *TreeNode, expectNumber int) [][]int {
	// write code here
	if root == nil {
		return nil
	}

	var (
		result [][]int
		path   []int
		dfs    func(node *TreeNode, cur int)
	)

	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && cur == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}
		dfs(node.Left, cur)
		dfs(node.Right, cur)
	}
	dfs(root,expectNumber)
	return result
}
