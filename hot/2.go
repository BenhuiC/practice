package hot

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	h := head
	y := 0
	for l1 != nil || l2 != nil || y != 0 {
		n := y
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}
		y = n / 10
		n = n % 10
		h.Next = &ListNode{Val: n}
		h = h.Next
	}

	return head.Next
}
