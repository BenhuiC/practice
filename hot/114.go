package hot

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var last *TreeNode
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last != nil {
			last.Right = t
		}
		last = t

		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		if t.Left != nil {
			stack = append(stack, t.Left)
		}
		t.Right = nil
		t.Left = nil
	}
}
