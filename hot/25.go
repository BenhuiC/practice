package hot

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}
	reverse := func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		l, r := head, head.Next
		l.Next = nil
		for r != nil {
			tmp := r.Next
			r.Next = l
			l = r
			r = tmp
		}
		return l
	}

	head = &ListNode{Next: head}
	h := head
	var next *ListNode
	for {
		tail := h.Next
		prev := h
		for i := 0; i < k && h != nil; i++ {
			h = h.Next
		}
		if h == nil {
			break
		}

		next = h.Next
		h.Next = nil

		// reverse
		prev.Next = reverse(tail)

		h = tail
		h.Next = next
	}

	return head.Next
}
