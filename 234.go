package partice

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
	}

	var h *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = h
		h = slow
		slow = tmp
	}

	for head != nil && h != nil {
		if head.Val != h.Val {
			return false
		}
		head = head.Next
		h = h.Next
	}
	return true
}
