package hot

func kthSmallest(root *TreeNode, k int) int {
	res := 0

	var find func(n *TreeNode)
	find = func(n *TreeNode) {
		if n == nil {
			return
		}
		find(n.Left)
		if k < 0 {
			return
		}
		k--
		if k == 0 {
			res = n.Val
		}
		find(n.Right)
	}

	find(root)
	return res
}
