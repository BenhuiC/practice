package link

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	m := make(map[*ListNode]struct{})
	for n := head; n != nil; n = n.Next {
		if _, ok := m[n]; ok {
			return n
		}
		m[n] = struct{}{}
	}

	return nil
}

/*
使用快慢指针
slow每次走一步，fast每次走两步
如有环，则两个指针肯定回相遇，此时快慢指针分别走了f，s步
加入链表长度为a+b，a为无环部分长度，b为有环部分长度
则
f=2*s
f=s+n*b (f比s多走了n圈)
可得
s=n*b
f=2*n*b

加入慢指针当前在环的入口处，它走过的长度为k
则k=a+n*b

由上可知快慢指针在第一次相遇时慢指针走了n*b的长度，而走到环的入口需要a+n*b的长度
所以此时慢指针只需要再走a步
此时可以把fast指针赋值为头节点，从这一步开始，slow和fast每次只走一步，当两个指针相遇时，即为入口节点
*/
func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for slow != nil && fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow == nil || fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
