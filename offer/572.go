package offer

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val && matchSub(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func matchSub(r, sub *TreeNode) bool {
	if r == nil && sub == nil {
		return true
	}
	if r == nil || sub == nil {
		return false
	}
	if r.Val != sub.Val {
		return false
	}
	return matchSub(r.Left, sub.Left) && matchSub(r.Right, sub.Right)
}
