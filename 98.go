package partice

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := true, true
	if root.Left != nil {
		ll := root.Left.Right
		for ll != nil && ll.Right != nil {
			ll = ll.Right
		}
		l = root.Val > root.Left.Val
		if ll != nil {
			l = l && ll.Val < root.Val
		}
	}
	if root.Right != nil {
		rr := root.Right.Left
		for rr != nil && rr.Left != nil {
			rr = rr.Left
		}
		r = root.Val < root.Right.Val
		if rr != nil {
			r = r && rr.Val > root.Val
		}
	}
	if !l || !r {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func isValidBST2(root *TreeNode) bool {
	return valid(root, ^int(^uint(0)>>1), int(^uint(0)>>1))
}

func valid(n *TreeNode, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Val <= min || n.Val >= max {
		return false
	}
	return valid(n.Left, min, n.Val) && valid(n.Right, n.Val, max)
}
