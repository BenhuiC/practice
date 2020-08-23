package partice

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

func inorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		n := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		root = n.Right
		result = append(result, n.Val)
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
