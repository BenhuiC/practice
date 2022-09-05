package partice

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	res := &TreeNode{}
	res.Val = postorder[n-1]
	idx := 0
	for ; idx < n; idx++ {
		if inorder[idx] == res.Val {
			break
		}
	}
	res.Left = buildTree(inorder[:idx], postorder[:idx])
	res.Right = buildTree(inorder[idx+1:], postorder[idx:n-1])

	return res
}
