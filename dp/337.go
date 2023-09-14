package dp

func rob337(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var robTree func(n *TreeNode) (rob, notRob int)
	robTree = func(n *TreeNode) (rob, notRob int) {
		if n == nil {
			return
		}
		rl, nrl := robTree(n.Left)
		ll, nll := robTree(n.Right)
		rob = nrl + nll + n.Val
		notRob = Max(rl, nrl) + Max(ll, nll)
		return
	}

	return Max(robTree(root))
}
