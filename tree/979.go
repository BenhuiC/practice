package tree

func distributeCoins(root *TreeNode) int {
	res := 0
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var dfs func(n *TreeNode) (int, int)
	dfs = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		leftVal, leftCount := dfs(n.Left)
		rightVal, rightCount := dfs(n.Right)
		res += abs(leftVal - leftCount)
		res += abs(rightVal - rightCount)

		return leftVal + rightVal + n.Val, leftCount + rightCount + 1
	}

	dfs(root)
	return res
}
