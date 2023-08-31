package partice

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isNodeSymmetric(root.Right, root.Left)
}

func isNodeSymmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		return isNodeSymmetric(l.Left, r.Right) && isNodeSymmetric(l.Right, r.Left)
	} else {
		return false
	}
}
