package recursion

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	h := &ListNode{Next: head}
	n1, n2 := h, h
	i := 1
	for tmp := h; tmp != nil; tmp = tmp.Next {
		if i == left {
			n1 = tmp
		}
		if i == right {
			n2 = tmp.Next.Next
			tmp.Next.Next = nil
			break
		}
		i++
	}
	tail := n1.Next
	n1.Next = reverseList2(n1.Next)
	tail.Next = n2
	if left == 1 {
		return n1.Next
	}
	return head
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var last, cur *ListNode = nil, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = last
		last = cur
		cur = tmp
	}
	return last
}
