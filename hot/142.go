package hot

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for slow != nil && fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			break
		}
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	if slow == nil || fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
