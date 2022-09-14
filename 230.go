package partice

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	if root == nil {
		return 0
	}
	var dfs func(n *TreeNode, k int) int
	dfs = func(n *TreeNode, k int) int {
		if n == nil {
			return 0
		}
		left := dfs(n.Left, k)
		if left >= k {
			return left
		}
		if left == k-1 {
			res = n.Val
			return left + 1
		}

		right := dfs(n.Right, k-left-1)

		return left + right + 1
	}
	_ = dfs(root, k)
	return res
}
