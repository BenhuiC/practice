package partice

func sortedArrayToBST(a []int) *TreeNode {
	n := len(a)
	if n == 0 {
		return nil
	}
	res := &TreeNode{}
	res.Val = a[n/2]
	res.Left = sortedArrayToBST(a[:n/2])
	res.Right = sortedArrayToBST(a[n/2+1:])
	return res
}
