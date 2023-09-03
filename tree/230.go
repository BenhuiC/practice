package tree

func kthSmallest(root *TreeNode, k int) int {
	idx := 0
	res := 0
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		if idx >= k {
			return
		}
		idx++
		if idx == k {
			res = n.Val
			return
		}
		dfs(n.Right)
	}
	dfs(root)

	return res
}
