package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	res := &ListNode{}
	for h := head; h != nil; {
		tmp := h.Next
		v := res
		for v.Next != nil && v.Next.Val < h.Val {
			v = v.Next
		}
		h.Next = v.Next
		v.Next = h
		h = tmp
	}

	return res.Next
}
