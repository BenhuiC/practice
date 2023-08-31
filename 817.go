package partice

func numComponents(head *ListNode, nums []int) int {
	if head==nil||len(nums)==0{
		return 0
	}
	mp:=make(map[int]bool,len(nums))
	for _,v:=range nums{
		mp[v]= true
	}
	var res817 int
	for p:=head;p!=nil;{
		if mp[p.Val]{
			res817=res817+1
			for p!=nil&&mp[p.Val]{
				p=p.Next
			}
		}else{
			p=p.Next
		}
	}

	return res817
}
