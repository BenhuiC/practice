package offer

func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{Next: head}
	for t := h; t != nil; t = t.Next {
		if t.Next != nil && t.Next.Val == val {
			var t2 *ListNode
			for t2 = t.Next; t2 != nil && t2.Val == val; t2 = t2.Next {
			}
			t.Next = t2
		}
	}

	return h.Next
}
