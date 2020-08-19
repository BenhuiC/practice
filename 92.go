package partice

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	result, mNode, nNode, vNode := &ListNode{Next: head}, &ListNode{}, &ListNode{}, &ListNode{}
	i := 0
	for p := result; p != nil; {
		if i+1 == m {
			mNode = p
		}
		if i >= m && i <= n {
			tmp := p.Next
			p.Next = vNode.Next
			vNode.Next = p
			p = tmp
			if i == n {
				nNode = p
				break
			}
		} else {
			p = p.Next
		}
		i++
	}
	mNode.Next = vNode

	return result.Next
}
