package partice

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}
