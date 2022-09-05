package partice

func longestUnivaluePath(root *TreeNode) int {
	return max(longestPath687(root))
}

func longestPath687(root *TreeNode) (withRoot, withoutRoot int) {
	if root == nil {
		return
	}
	var lw, lwo, rw, rwo int
	if root.Left != nil {
		lw, lwo = longestPath687(root.Left)
		if root.Left.Val == root.Val {
			withRoot = max(withRoot, lw+1)
			withoutRoot += lw + 1
		}
	}
	if root.Right != nil {
		rw, rwo = longestPath687(root.Right)
		if root.Right.Val == root.Val {
			withRoot = max(withRoot, rw+1)
			withoutRoot += rw + 1
		}
	}
	withoutRoot = max(withoutRoot, max(lwo, rwo))
	return
}
