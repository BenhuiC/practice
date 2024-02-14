package hot

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt

	max := func(x ...int) int {
		r := math.MinInt
		for _, v := range x {
			if v > r {
				r = v
			}
		}
		return r
	}

	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		left, right := dfs(n.Left), dfs(n.Right)
		res = max(res, n.Val, left+n.Val, right+n.Val, left+right+n.Val)
		return max(left+n.Val, right+n.Val, n.Val)
	}

	dfs(root)
	return res
}
