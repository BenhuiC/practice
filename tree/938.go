package tree

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var res int
	if root.Val <= high && root.Val >= low {
		res = root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		res = rangeSumBST(root.Left, low, high)
	} else {
		res = rangeSumBST(root.Right, low, high)
	}
	return res
}
