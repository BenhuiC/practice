package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	val := nums[0]
	for i, v := range nums {
		if v > val {
			val = v
			idx = i
		}
	}
	root := &TreeNode{Val: val}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}
