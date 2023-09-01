package tree

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	idx := 0
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			idx = i
			break
		}
	}
	root.Left = buildTree2(inorder[:idx], postorder[:idx])
	root.Right = buildTree2(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}
