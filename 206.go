package partice

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := &ListNode{}
	for head != nil {
		tmp := head.Next
		head.Next = result.Next
		result.Next = head
		head = tmp
	}
	return result.Next
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
