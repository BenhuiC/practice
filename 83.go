package partice

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	n := head
	for n != nil {
		p := n.Next
		for p != nil && n.Val == p.Val {
			p = p.Next
		}
		n.Next = p
		n = n.Next
	}

	return head
}
