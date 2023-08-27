package link

import "fmt"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}

	h := &ListNode{Next: head}
	pNode, right, left, next := h, h.Next, h, h
	idx := 1
	for t := h; t != nil; t = t.Next {
		if idx%k == 0 {
			if t.Next == nil {
				break
			}
			left = t.Next
			next = left.Next
			left.Next = nil
			pNode.Next = reverseList(pNode.Next)
			right.Next = next
			t = right
			pNode = right
			right = next
			printList(h.Next)
			fmt.Println("pnode ", pNode)
			fmt.Println("right ", right)
			idx = 1
		}
		idx++
	}

	return h.Next
}
