package tree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs := make([]int, 0)
	var dfs1 func(n *TreeNode)
	dfs1 = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			leafs = append(leafs, n.Val)
		}
		dfs1(n.Left)
		dfs1(n.Right)
	}
	dfs1(root1)

	idx := 0
	var dfs2 func(n *TreeNode) bool
	dfs2 = func(n *TreeNode) bool {
		if n == nil {
			return true
		}
		if n.Left == nil && n.Right == nil {
			if idx >= len(leafs) || n.Val != leafs[idx] {
				return false
			}
			idx++
		}
		return dfs2(n.Left) && dfs2(n.Right)
	}
	return dfs2(root2) && idx == len(leafs)
}
