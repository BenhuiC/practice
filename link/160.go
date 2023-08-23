package link

/*
思路
1. 暴力法：用map记录链表a的节点，再遍历b的节点。时间复杂度O(m+n),空间复杂度O(m+n)；
2. 反转链表：反转a和b，再从头开始找到最后一个相等的节点。时间复杂度O(m+n),空间复杂度O(1)；链表必须保持其原始结构，此方案pass
3. 较长链表长度为m，短链表长度为n。则长链表需要从m-n个节点开始依次与短链表对比是否相同
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for a := headA; a != nil; a = a.Next {
		m[a] = struct{}{}
	}
	for b := headB; b != nil; b = b.Next {
		if _, ok := m[b]; ok {
			return b
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m, n := headA, headB
	// 当m==nil&&n==nil时退出循环
	for m != n {
		if m == nil {
			m = headB
		} else {
			m = m.Next
		}
		if n == nil {
			n = headA
		} else {
			n = n.Next
		}
	}
	return m
}
