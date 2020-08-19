package partice

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := &ListNode{Next: head}
	head = dump
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal := head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return dump.Next
}
