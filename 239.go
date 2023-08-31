package partice

func maxSlidingWindow(nums []int, k int) []int {
	res239:=make([]int,0,len(nums)-k+1)
	maxq:=make([]int,0)
	for i:=0;i<k;i++{
		j:=len(maxq)-1
		for ;j>=0&&nums[maxq[j]]<nums[i];j--{
			maxq=maxq[:len(maxq)-1]
		}
		maxq = append(maxq, i)
	}
	res239 = append(res239, nums[maxq[0]])
	
	for i:=k;i<len(nums);i++{
		for j:=len(maxq)-1;j>=0&&nums[maxq[j]]<nums[i];j--{
			maxq=maxq[:len(maxq)-1]
		}
		maxq = append(maxq, i)
		if nums[maxq[len(maxq)-1]]>nums[i-k]{
			res239 = append(res239, nums[maxq[0]])
			continue
		}

		for t:=0;t<len(maxq)&&maxq[t]<=i-k;t++{
			maxq=maxq[1:]
		}

		res239 = append(res239, nums[maxq[0]])
	}
	
	return res239
}
