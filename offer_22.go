package partice

func getKthFromEnd(head *ListNode, k int) *ListNode {
	l:=1
	for v:=head;v!=nil;v=v.Next{
		l++
	}
	t:=l-k
	resOffer22:=head
	for i:=1;resOffer22!=nil&&t!=i;resOffer22=resOffer22.Next{
		i++
	}
	return resOffer22
}
