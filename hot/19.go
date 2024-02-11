package hot

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Next: head}
	tail := head
	for i := 0; i < n; i++ {
		if tail == nil {
			return nil
		}
		tail = tail.Next
	}
	prew := head
	for tail.Next != nil {
		tail = tail.Next
		prew = prew.Next
	}

	prew.Next = prew.Next.Next
	return head.Next
}
