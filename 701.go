package partice

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	n := root
	for {
		if val > n.Val {
			if n.Right == nil {
				n.Right = &TreeNode{Val: val}
				break
			}
			n = n.Right
		} else {
			if n.Left == nil {
				n.Left = &TreeNode{Val: val}
				break
			}
			n = n.Left
		}
	}

	return root
}
