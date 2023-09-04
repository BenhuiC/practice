package tree

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode

	var f func(root, p, q *TreeNode) int
	f = func(root, p, q *TreeNode) int {
		if res != nil {
			return 0
		}
		if root == nil {
			return 0
		}
		l := f(root.Left, p, q)
		r := f(root.Right, p, q)
		if l == 1 && r == 1 {
			res = root
			return 0
		}
		if root.Val == p.Val || root.Val == q.Val {
			if l == 1 || r == 1 {
				res = root
				return 0
			}
			return l + r + 1
		}
		if l != 0 {
			return l
		}
		return r
	}
	f(root, p, q)
	return res
}
