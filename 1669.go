package partice

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head:=&ListNode{Next: list1}
	var st,ed *ListNode
	var i int
	for n:=head;n!=nil;n=n.Next{
		if i==a{
			st=n
		}
		if i==b+1{
			ed=n
		}
		i++
	}
	if list2==nil{
		st.Next=ed.Next
		return head.Next
	}

	var ed2 *ListNode
	for ed2=list2;ed2!=nil&&ed2.Next!=nil;ed2=ed2.Next{}
	st.Next=list2
	ed2.Next=ed.Next
	ed.Next=nil


	return head.Next
}
