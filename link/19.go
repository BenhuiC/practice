package link

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{
		Next: head,
	}
	l1 := h
	for i := 0; i < n; i++ {
		l1 = l1.Next
	}
	l2 := h
	for l1 != nil && l1.Next != nil {
		l1 = l1.Next
		l2 = l2.Next
	}
	l2.Next = l2.Next.Next
	return h.Next
}
