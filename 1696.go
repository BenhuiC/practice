package partice

func maxResult(nums []int, k int) int {
	// 单调递减，队首最大
	maxQ:=make([]int,0,len(nums))
	for i:=range nums{
		// 队列为空，进队列
		if len(maxQ)==0{
			maxQ = append(maxQ, i)
			continue
		}
		// 队首要在i-k内
		j:=0
		for maxQ[j]<i-k{
			j++
		}
		maxQ=maxQ[j:]
		// 当前值=nums[i]+最大值
		nums[i]+=nums[maxQ[0]]
		// 将当前值入队
		for len(maxQ)>0&&nums[maxQ[len(maxQ)-1]]<nums[i]{
			maxQ=maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)
	}

	return nums[len(nums)-1]
}