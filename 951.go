package partice

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	if root1.Left == nil {
		if root2.Left == nil {
			return flipEquiv(root1.Right, root2.Right)
		}
		return root2.Right == nil && flipEquiv(root1.Right, root2.Left)
	} else {
		if root2.Left == nil {
			return root1.Right == nil && flipEquiv(root1.Left, root2.Right)
		}
		if root1.Left.Val == root2.Left.Val {
			return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
		} else {
			return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
		}
	}
}
