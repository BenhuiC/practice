package partice

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	res := &TreeNode{Val: nums[maxIndex]}
	res.Left = constructMaximumBinaryTree(nums[:maxIndex])
	res.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return res
}

// todo 单调栈实现
