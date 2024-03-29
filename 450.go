package partice

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		return delete450(root)
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func delete450(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return nil
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {
		if root.Left.Right == nil {
			root.Val = root.Left.Val
			root.Left = delete450(root.Left)
		} else {
			pre, n := root.Left, root.Left.Right
			for n != nil && n.Right != nil {
				pre = n
				n = n.Right
			}
			root.Val = n.Val
			pre.Right = delete450(n)
		}
		return root
	}
}
