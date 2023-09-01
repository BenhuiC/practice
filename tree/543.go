package tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var deep func(r *TreeNode) int
	deep = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := deep(r.Left)
		right := deep(r.Right)
		res = Max(res, left+right)
		return Max(left, right) + 1
	}
	_ = deep(root)
	return res
}
