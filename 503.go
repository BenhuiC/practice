package partice

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stk := make([]int, 0)
	for i := range res {
		res[i] = -1
	}
	for i := 0; i < 2*n-1; i++ {
		relIdx := i % n
		if len(stk) == 0 || nums[stk[len(stk)-1]] >= nums[relIdx] {
			stk = append(stk, relIdx)
			continue
		}
		for len(stk) > 0 && nums[stk[len(stk)-1]] < nums[relIdx] {
			res[stk[len(stk)-1]] = nums[relIdx]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, relIdx)
	}

	return res
}
