package partice

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return calculate(l1, l2, 0)
}

func calculate(l1, l2 *ListNode, n int) *ListNode {
	if l1 == nil {
		return calculate2(l2, n)
	}
	if l2 == nil {
		return calculate2(l1, n)
	}
	t := l1.Val + l2.Val + n
	if t > 9 {
		t = t - 10
		n = 1
	} else {
		n = 0
	}
	l := &ListNode{
		Val:  t,
		Next: calculate(l1.Next, l2.Next, n),
	}
	return l
}

func calculate2(l1 *ListNode, n int) *ListNode {
	if l1 == nil && n == 0 {
		return nil
	}
	if l1 == nil {
		return &ListNode{
			Val:  n,
			Next: nil,
		}
	}
	if n == 0 {
		return l1
	}
	t := l1.Val + n
	if t > 9 {
		t = t - 10
		n = 1
	} else {
		n = 0
	}
	l := &ListNode{
		Val:  t,
		Next: calculate2(l1.Next, n),
	}
	return l
}
