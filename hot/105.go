package hot

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}
	for i, v := range inorder {
		if v != preorder[0] {
			continue
		}
		root.Left = buildTree(preorder[1:i+1], inorder[:i])
		root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	}

	return root
}
