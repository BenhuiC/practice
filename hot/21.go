package hot

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	h := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			h.Next = list1
			list1 = list1.Next
		} else {
			h.Next = list2
			list2 = list2.Next
		}
		h = h.Next
	}
	if list1 != nil {
		h.Next = list1
	}
	if list2 != nil {
		h.Next = list2
	}

	return head.Next
}
