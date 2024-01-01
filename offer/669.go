package offer

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val >= low && root.Val <= high {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
	} else if root.Val < low {
		return trimBST(root.Right, low, high)
	} else {
		return trimBST(root.Left, low, high)
	}
	return root
}
