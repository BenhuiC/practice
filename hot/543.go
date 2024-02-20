package hot

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	var deep func(t *TreeNode) int
	deep = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		left, right := deep(t.Left), deep(t.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}

	deep(root)
	return res
}
