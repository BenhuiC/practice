package partice

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	h := &TreeNode{Val: preorder[0]}
	rIdx := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rIdx = i
			break
		}
	}
	h.Left = buildTree(preorder[1:rIdx+1], inorder[:rIdx])
	h.Right = buildTree(preorder[rIdx+1:], inorder[rIdx+1:])
	return h
}
