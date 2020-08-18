package partice

func isBalanced(root *TreeNode) bool {
	_, b := balance(root)
	return b
}

func balance(n *TreeNode) (int, bool) {
	var d int
	if n == nil {
		return 0, true
	}
	left, lb := balance(n.Left)
	right, rb := balance(n.Right)

	if !lb || !rb {
		return 0, false
	}

	if left-right > 1 || right-left > 1 {
		return 0, false
	}

	if left > right {
		d = left + 1
	} else {
		d = right + 1
	}

	return d, true
}
