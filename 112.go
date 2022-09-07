package partice

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	//if targetSum < 0 {
	//	return false
	//}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	targetSum -= root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
