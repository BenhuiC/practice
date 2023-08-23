package link

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	res := list1
	l, h := list1, list2
	if list1.Val > list2.Val {
		l, h = list2, list1
		res = list2
	}
	for l != nil && h != nil {
		if l.Next == nil {
			l.Next = h
			break
		}
		if l.Next.Val > h.Val {
			tmp := h.Next
			h.Next = l.Next
			l.Next = h
			l = h
			h = tmp
		} else {
			l = l.Next
		}
	}
	return res
}
