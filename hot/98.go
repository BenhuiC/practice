package hot

import "math"

func isValidBST(root *TreeNode) bool {
	var valid func(n *TreeNode, min, max int) bool
	valid = func(n *TreeNode, min, max int) bool {
		if n == nil {
			return true
		}
		if n.Val >= min || n.Val <= max {
			return false
		}
		return valid(n.Left, n.Val, max) && valid(n.Right, min, n.Val)
	}

	return valid(root, math.MaxInt, math.MinInt)
}
