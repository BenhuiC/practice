package offer

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil && root.Right == nil {
		return -1
	}
	left, right := -1, -1
	if root.Left.Val == root.Val && root.Right.Val == root.Val {
		left = findSecondMinimumValue(root.Left)
		right = findSecondMinimumValue(root.Right)
	} else if root.Left.Val == root.Val {
		left = findSecondMinimumValue(root.Left)
		right = root.Right.Val
	} else if root.Right.Val == root.Val {
		left = root.Left.Val
		right = findSecondMinimumValue(root.Right)
	}

	if left == -1 {
		return right
	} else if right == -1 {
		return left
	}
	return min(left, right)
}
