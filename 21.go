package partice

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{}
	p := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp := l1.Next
			p.Next = l1
			p = l1
			l1 = tmp
		} else {
			tmp := l2.Next
			p.Next = l2
			p = l2
			l2 = tmp
		}
	}
	if l1 == nil {
		p.Next = l2
	}
	if l2 == nil {
		p.Next = l1
	}
	return result.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val < l2.Val {
		result = l1
		l1.Next = mergeTwoLists2(l1.Next, l2)
	} else {
		result = l2
		l2.Next = mergeTwoLists2(l1, l2.Next)
	}

	return result
}
