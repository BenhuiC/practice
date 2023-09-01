package tree

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rval := preorder[0]
	root := &TreeNode{Val: rval}
	if len(preorder) == 1 {
		return root
	}
	rightVal := postorder[len(postorder)-2]
	rightIdx := 0
	for i, v := range preorder {
		if v == rightVal {
			rightIdx = i
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:rightIdx], postorder[:rightIdx-1])
	root.Right = constructFromPrePost(preorder[rightIdx:], postorder[rightIdx-1:len(postorder)-1])

	return root
}
