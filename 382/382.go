package _382

import "math/rand"

/**
* Definition for singly-linked list.
*
*/
type ListNode struct {
    Val int
    Next *ListNode
}


type Solution struct {
	*ListNode
	Len int64
}


/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	s:=Solution{
		ListNode: head,
		Len:      0,
	}
	var i int64
	for p:=head;p!=nil;i++{
		p=p.Next
	}
	s.Len=i
	return s
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	if this.ListNode==nil{
		return 0
	}
	r:=rand.Int63n(this.Len)
	var i int64
	for p:=this.ListNode;p!=nil;i++{
		if i==r{
			return p.Val
		}
		p=p.Next
	}
	return 0
}
