package partice

func levelOrderBottom(root *TreeNode) [][]int {
	vec := make([][]int, 0)
	if root == nil {
		return vec
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		l := make([]int, 0)
		for _, v := range queue {
			l = append(l, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[n:]
		tmp := [][]int{l}
		vec = append(tmp, vec...)
	}

	return vec
}
