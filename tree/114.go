package tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var last *TreeNode
	for len(stack) != 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left, right := t.Left, t.Right
		t.Left, t.Right = nil, nil
		if last != nil {
			last.Right = t
		}
		last = t

		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
	}

}
