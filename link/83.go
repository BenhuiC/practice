package link

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for h := head; h.Next != nil; {
		if h.Next.Val == h.Val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}

	return head
}
