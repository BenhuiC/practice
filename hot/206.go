package hot

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l, r := head, head.Next
	l.Next = nil
	for r != nil {
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	return l
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
