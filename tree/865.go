package tree

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var res *TreeNode
	var maxDep int

	var dfs func(n *TreeNode, d int) int
	dfs = func(n *TreeNode, d int) int {
		if n == nil {
			return d
		}
		left := dfs(n.Left, d+1)
		right := dfs(n.Right, d+1)

		if left == right && left >= maxDep {
			res = n
			maxDep = left
		}

		return max(left, right)
	}

	dfs(root, 0)

	return res
}
