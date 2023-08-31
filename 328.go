package partice

func oddEvenList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	aHead:=head
	bHead:=head.Next
	i:=1
	var p *ListNode
	for p=bHead.Next;p!=nil;{
		if i%2==1{
			bHead.Next=p.Next
			p.Next=aHead.Next
			aHead.Next=p
			aHead=p
			p=bHead.Next
		}else{
			bHead=p
			p=p.Next
		}
		i++
	}
	return head
}
