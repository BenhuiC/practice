package partice

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummp := &ListNode{Val: 0, Next: head}
	// 计算前缀和，即前缀和之间的数字和为0
	prefixSum := 0
	seen := make(map[int]*ListNode)
	for n := dummp; n != nil; n = n.Next {
		prefixSum += n.Val
		seen[prefixSum] = n
	}

	prefixSum = 0
	for n := dummp; n != nil; n = n.Next {
		prefixSum += n.Val
		n.Next = seen[prefixSum].Next
	}

	return dummp.Next
}

/*
(x,y]
1 1 2 3 -3 -2 -1 1
1 2 4 7  4  2  1 2
*/
