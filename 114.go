package partice

func flatten(root *TreeNode) {
	fl(root)
}

func fl(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	l := fl(root.Left)
	r := fl(root.Right)
	root.Left = nil
	if l == nil {
		root.Right = r
		return root
	}
	h := l
	for ; h.Right != nil; h = h.Right {
	}
	h.Right = r
	root.Right = l
	return root
}
