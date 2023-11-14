package offer

// o ho,timeout
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var addTwo func(n1, n2 *ListNode) (n *ListNode, b int)
	addTwo = func(n1, n2 *ListNode) (n *ListNode, b int) {
		if n1 == nil && n2 == nil {
			return
		}
		var v, sub int
		if n1 == nil {
			n2.Next, sub = addTwo(nil, n2.Next)
			v = n2.Val + sub
			b = v / 10
			n2.Val = v % 10
			return n2, b
		}
		if n2 == nil {
			n1.Next, sub = addTwo(n1.Next, nil)
			v = n1.Val + sub
			b = v / 10
			n1.Val = v % 10
			return n1, b
		}
		if n1.Next == nil && n2.Next == nil {
			v = n1.Val + n2.Val
			b = v / 10
			n1.Val = v % 10
			return n1, b
		}
		if n1.Next == nil {
			n1.Next, sub = addTwo(n1, n2.Next)
			v = n2.Val + sub
			b = v / 10
			n1.Val = v % 10
			return n1, b
		}
		if n2.Next == nil {
			n1.Next, sub = addTwo(n1.Next, n2)
			v = n1.Val + sub
			b = v / 10
			n1.Val = v % 10
			return n1, b
		}
		return addTwo(n1.Next, n2.Next)
	}
	res, sub := addTwo(l1, l2)
	if sub != 0 {
		res = &ListNode{Val: 1, Next: res}
	}
	return res
}
