package hot

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
