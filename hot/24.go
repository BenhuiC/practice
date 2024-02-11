package hot

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{Next: head}
	h := head
	for h.Next != nil {
		l := h.Next
		r := l.Next
		if r == nil {
			break
		}
		t := r.Next
		h.Next = r
		r.Next = l
		l.Next = t
		h = l
	}

	return head.Next
}
