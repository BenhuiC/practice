package partice

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	rev := &ListNode{}
	i := 0
	for p := head; p != nil; {
		n := &ListNode{Val: p.Val}
		n.Next = rev.Next
		rev.Next = n
		i++
		p = p.Next
	}
	if i%2 == 0 {
		cut(head, i/2)
		cut(rev.Next, i/2)
	} else {
		cut(head, i/2+1)
		cut(rev.Next, i/2)
	}
	merge(head, rev.Next)
}

func cut(l *ListNode, n int) {
	i := 0
	for r := l; r != nil; r = r.Next {
		i++
		if i == n {
			r.Next = nil
			return
		}
	}
}

func merge(l1, l2 *ListNode) {
	for p, q := l1, l2; q != nil && p != nil; {
		tmp := p.Next
		tmp2 := q.Next
		p.Next = q
		p.Next.Next = tmp
		p = tmp
		q = tmp2
	}
}
