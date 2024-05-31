package tree

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	res := 0
	var dfs func(n *TreeNode) (int, int)
	dfs = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		lmin, lmax := n.Val, n.Val
		if n.Left != nil {
			lmin, lmax = dfs(n.Left)
		}
		rmin, rmax := n.Val, n.Val
		if n.Right != nil {
			rmin, rmax = dfs(n.Right)
		}
		mi, mx := min(lmin, rmin), max(lmax, rmax)
		res = max(abs(n.Val-mi), abs(n.Val-mx), res)
		if n.Val < mi {
			mi = n.Val
		}
		if n.Val > mx {
			mx = n.Val
		}
		return mi, mx
	}
	dfs(root)

	return res
}
