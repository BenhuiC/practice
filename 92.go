package partice

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	dump, pre, nNode, vNode := &ListNode{Next: head}, &ListNode{}, &ListNode{}, &ListNode{}
	i := 0
	for p := dump; p != nil; {
		if i+1 == m {
			nNode = p.Next
			pre = p
			p = p.Next
		} else if i >= m && i <= n {
			tmp := p.Next
			p.Next = vNode.Next
			vNode.Next = p
			p = tmp
			if i == n {
				nNode.Next = p
				break
			}
		} else {
			p = p.Next
		}
		i++
	}
	pre.Next = vNode.Next

	return dump.Next
}
