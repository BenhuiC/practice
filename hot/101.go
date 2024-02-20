package hot

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isSymmNode func(l, r *TreeNode) bool
	isSymmNode = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		} else if l != nil && r != nil {
			if l.Val != r.Val {
				return false
			}
			return isSymmNode(l.Left, r.Right) && isSymmNode(l.Right, r.Left)
		} else {
			return false
		}
	}
	return isSymmNode(root.Left, root.Right)
}
