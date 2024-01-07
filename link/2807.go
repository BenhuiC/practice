package link

func insertGreatestCommonDivisors(head *ListNode) *ListNode {

	for cur := head; cur != nil; {
		if cur.Next == nil {
			break
		}
		i, j := cur.Val, cur.Next.Val
		n := &ListNode{Val: gcd(i, j), Next: cur.Next}
		cur.Next = n
		cur = n.Next
	}

	return head
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else if a < b {
			b -= a
		} else {
			return a
		}
	}
	return a
}
