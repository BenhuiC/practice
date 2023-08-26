package link

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := &ListNode{}
	tmp := l
	h := &ListNode{
		Next: head,
	}
	for i := h; i != nil; {
		if i.Next != nil && i.Next.Val < x {
			tmp.Next = i.Next
			tmp = tmp.Next
			i.Next = tmp.Next
		} else {
			i = i.Next
		}
	}
	tmp.Next = h.Next

	return l.Next
}
