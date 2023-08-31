package partice

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	lmax := maxNode(root.Left)
	rmin := minNode(root.Right)
	if lmax != nil && lmax.Val > root.Val && rmin != nil && rmin.Val < root.Val {
		lmax.Val, rmin.Val = rmin.Val, lmax.Val
	}
	if lmax != nil && lmax.Val > root.Val {
		lmax.Val, root.Val = root.Val, lmax.Val
		return
	}
	if rmin != nil && rmin.Val < root.Val {
		rmin.Val, root.Val = root.Val, rmin.Val
		return
	}
	recoverTree(root.Left)
	recoverTree(root.Right)
}

func maxNode(n *TreeNode) *TreeNode {
	if n == nil {
		return n
	}
	max := n
	if ln := maxNode(n.Left); ln != nil && ln.Val > max.Val {
		max = ln
	}
	if rn := maxNode(n.Right); rn != nil && rn.Val > max.Val {
		max = rn
	}
	return max
}

func minNode(n *TreeNode) *TreeNode {
	if n == nil {
		return n
	}
	min := n
	if ln := minNode(n.Left); ln != nil && ln.Val < min.Val {
		min = ln
	}
	if rn := minNode(n.Right); rn != nil && rn.Val < min.Val {
		min = rn
	}
	return min
}
