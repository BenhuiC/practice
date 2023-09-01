package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	root := &TreeNode{Val: val}
	inorderIdx := 0
	for i, v := range inorder {
		if val == v {
			inorderIdx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:inorderIdx+1], inorder[:inorderIdx])
	root.Right = buildTree(preorder[inorderIdx+1:], inorder[inorderIdx+1:])

	return root
}
