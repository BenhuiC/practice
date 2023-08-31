package partice

func rotateRight(head *ListNode, k int) *ListNode {
	if k==0{
		return head
	}
	var l int
	var last *ListNode
	for last=head;last!=nil&&last.Next!=nil;last=last.Next{
		l++
	}
	l=l+1

	var realMov int
	if realMov=k%l;realMov==0{
		return head
	}
	realMov=l-realMov
	i,tmp:=1,head
	for ;i<realMov;i++{
		tmp=tmp.Next
	}
	last.Next=head
	head=tmp.Next
	tmp.Next=nil


	return head
}