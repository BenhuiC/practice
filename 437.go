package partice

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res437 := numSumTarget(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)

	return res437
}

func numSumTarget(root *TreeNode, tar int) int {
	res := 0
	if root == nil {
		return 0
	}
	if root.Val == tar {
		res = 1
	}
	last := tar - root.Val
	res += numSumTarget(root.Left, last)
	res += numSumTarget(root.Right, last)
	return res
}
