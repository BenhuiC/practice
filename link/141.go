package link

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	n, next := head, head
	for n != nil && next != nil && next.Next != nil {
		n = n.Next
		next = next.Next.Next
		if n == next {
			return true
		}
	}

	return false
}
