package offer

import "math"

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}
	left := math.MaxInt
	right := math.MaxInt

	if root.Left != nil {
		left = root.Val - root.Left.Val
	}
	if root.Right != nil {
		right = root.Right.Val - root.Val
	}
	cur := min(left, right)

	leftM := minDiffInBST(root.Left)
	rightM := minDiffInBST(root.Right)
	return min(cur, min(leftM, rightM))
}
