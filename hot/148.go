package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return mergeSort(head)
}

func findMid(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge2(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	r := &ListNode{}
	t := r
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp := l1.Next
			t.Next = l1
			t = l1
			l1 = tmp
		} else {
			tmp := l2.Next
			t.Next = l2
			t = l2
			l2 = tmp
		}
	}
	if l1 == nil {
		t.Next = l2
	}
	if l2 == nil {
		t.Next = l1
	}

	return r.Next
}

func mergeSort(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	mid := findMid(l)
	tail := mid.Next
	mid.Next = nil
	left := mergeSort(l)
	right := mergeSort(tail)
	result := merge2(left, right)
	return result
}
