package recursion

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := reverseList(head.Next)
	if head.Next != nil {
		head.Next.Next = head
		head.Next = nil
	}
	return l
}
