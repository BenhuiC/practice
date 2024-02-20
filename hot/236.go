package hot

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode

	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if res != nil || n == nil {
			return 0
		}
		l, r := dfs(n.Left), dfs(n.Right)
		if l == 1 && r == 1 {
			res = n
			return 0
		}
		if p.Val == n.Val || q.Val == n.Val {
			if l == 1 || r == 1 {
				res = n
				return 0
			}
			return l + r + 1
		}
		if l != 0 {
			return l
		}
		return r
	}

	dfs(root)

	return res
}
