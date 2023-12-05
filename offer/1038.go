package offer

func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	var gstSum func(n *TreeNode)
	gstSum = func(n *TreeNode) {
		if n == nil {
			return
		}
		gstSum(n.Right)
		sum += n.Val
		n.Val = sum
		gstSum(n.Left)
	}
	gstSum(root)
	return root
}
