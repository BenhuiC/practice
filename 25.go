package partice

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	h := &ListNode{Next: head}
	l, f := h, h
	var i int
	for i = 0; i <= k && f != nil; i++ {
		f = f.Next
	}
	if i < k+1 {
		return head
	}
	for {
		h, e := reverseL(l.Next, f)
		l.Next = h
		e.Next = f
		l = e
		var j int
		for j = 0; j < k && f != nil; j++ {
			f = f.Next
		}
		if j < k {
			break
		}
	}

	return h.Next
}

func reverseL(begin, end *ListNode) (h, e *ListNode) {
	h = begin
	e = begin
	t := h.Next
	h.Next = nil
	for t != end {
		tmp := t.Next
		t.Next = h
		h = t
		t = tmp
	}
	return
}
