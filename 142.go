package partice

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := make(map[*ListNode]bool)
	for p := head; p != nil; p = p.Next {
		if _, ok := m[p]; ok {
			return p
		}
		m[p] = true
	}
	return nil
}
