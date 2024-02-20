package hot

func maxSubArray(nums []int) int {
	res := nums[0]
	n := len(nums)
	if n < 1 {
		return 0
	}
	curMax := nums[0]
	for i := 1; i < n; i++ {
		if curMax > 0 {
			curMax += nums[i]
		} else {
			curMax = nums[i]
		}
		if curMax > res {
			res = curMax
		}
	}

	return res
}
