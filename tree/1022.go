package tree

func sumRootToLeaf(root *TreeNode) int {
	res := 0
	var dfs func(n *TreeNode, cur int)
	dfs = func(n *TreeNode, cur int) {
		if n == nil {
			return
		}
		cur = cur*2 + n.Val
		if n.Left == nil && n.Right == nil {
			res += cur
			return
		}
		if n.Left != nil {
			dfs(n.Left, cur)
		}
		if n.Right != nil {
			dfs(n.Right, cur)
		}
	}

	dfs(root, 0)
	return res
}
