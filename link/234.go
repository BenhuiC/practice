package link

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	res := true
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	left := head
	var right *ListNode
	if fast == nil {
		right = reverseList2(slow)
	} else {
		right = reverseList2(slow.Next)
	}

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return res
}
