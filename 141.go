package partice

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	m := make(map[*ListNode]bool)
	for p := head; p != nil; p = p.Next {
		if m[p] {
			return true
		}
		m[p] = true
	}
	return false
}
