package partice

func splitListToParts(head *ListNode, k int) []*ListNode {
	if k<=1{
		return []*ListNode{head}
	}
	l:=getLen(head)
	x,y:=l/k,l%k
	n:=head
	res725:=make([]*ListNode,k)
	for i:=0;i<k&&n!=nil;i++{
		tmpLen:=x
		if i<y{
			tmpLen++
		}
		tmpI,tmpNode:=1,n
		for ;tmpI<tmpLen;tmpI++{
			tmpNode=tmpNode.Next
		}

		res725[i]=n
		n=tmpNode.Next
		tmpNode.Next=nil
	}

	return res725
}


func getLen(n *ListNode) int{
	i:=0
	for ;n!=nil;n=n.Next{
		i++
	}
	return i
}
