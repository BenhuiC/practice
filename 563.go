package partice

func findTilt(root *TreeNode) int {
	var res int
	var bac func(n *TreeNode) int
	bac = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := bac(root.Left), bac(root.Right)
		res += abs(l - r)
		return l + r + n.Val
	}
	bac(root)
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
