package offer

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	var dfs func(t *TreeNode) bool
	dfs = func(t *TreeNode) bool {
		if t == nil {
			return false
		}
		if _, ok := m[k-t.Val]; ok {
			return true
		}
		m[t.Val] = struct{}{}
		return dfs(t.Left) || dfs(t.Right)
	}
	return dfs(root)
}
