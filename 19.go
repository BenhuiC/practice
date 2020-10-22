package partice

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	h := &ListNode{Next: head}
	remove(h, n)
	return h.Next
}

func remove(node *ListNode, target int) int {
	if node == nil {
		return 0
	}
	n := remove(node.Next, target)
	if n == target {
		tmp := node.Next
		node.Next = tmp.Next
	}
	return n + 1
}
