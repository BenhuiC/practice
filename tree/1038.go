package tree

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(n *TreeNode, base int) int
	dfs = func(n *TreeNode, base int) int {
		if n == nil {
			return 0
		}
		v := n.Val
		rsum := dfs(n.Right, base)
		n.Val = n.Val + base + rsum
		lsum := dfs(n.Left, n.Val)
		return v + rsum + lsum
	}
	dfs(root, 0)

	return root
}

func bstToGst2(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}
