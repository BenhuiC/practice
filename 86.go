package partice

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	lower, higher := &ListNode{}, &ListNode{}
	l, h := lower, higher
	for head != nil {
		if head.Val < x {
			tmp := head.Next
			l.Next = head
			l = l.Next
			l.Next = nil
			head = tmp
		} else {
			tmp := head.Next
			h.Next = head
			head = head.Next
			h = h.Next
			h.Next = nil
			head = tmp
		}
	}

	l.Next = higher.Next

	return lower.Next
}
