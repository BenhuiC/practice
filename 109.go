package partice

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	ary := make([]int, 0)
	for head != nil {
		ary = append(ary, head.Val)
		head = head.Next
	}
	res := sortedArrayToBST(ary)

	return res
}

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
