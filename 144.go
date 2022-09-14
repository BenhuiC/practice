package partice

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(r *TreeNode)
	preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		res = append(res, r.Val)
		preorder(r.Left)
		preorder(r.Right)
		return
	}
	preorder(root)

	return res
}
