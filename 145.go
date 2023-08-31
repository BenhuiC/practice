package partice

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(r *TreeNode)
	postorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		postorder(r.Left)
		postorder(r.Right)
		res = append(res, r.Val)
		return
	}
	postorder(root)

	return res
}
