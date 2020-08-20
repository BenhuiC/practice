package partice

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	st := make([]*TreeNode, 0)
	for root != nil || len(st) > 0 {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		n := st[len(st)-1]
		st = st[:len(st)-1]
		result = append(result, n.Val)
		root = n.Right
	}

	return result
}
