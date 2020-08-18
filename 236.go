package partice

var v *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if sonOf(p, q) {
		return p
	}
	if sonOf(q, p) {
		return q
	}

	h(root, p, q)
	return v
}

func h(r, p, q *TreeNode) int {
	if r == nil {
		return -1
	}
	if r.Val == p.Val || r.Val == q.Val {
		return 1
	}
	left := h(r.Left, p, q)
	right := h(r.Right, p, q)
	if left > 0 && right > 0 {
		v = r
	}
	if left > 0 {
		return left
	} else {
		return right
	}
}

func sonOf(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return true
	}
	return sonOf(p.Left, q) || sonOf(p.Right, q)
}
