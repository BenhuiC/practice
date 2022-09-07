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
