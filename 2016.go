package partice

func maximumDifference(nums []int) int {
	res := 0
	if len(nums) < 2 {
		return 0
	}
	n := len(nums)
	maxV := nums[n-1]
	res = -1
	for i := n - 2; i >= 0; i-- {
		if maxV > nums[i] {
			res = max(res, maxV-nums[i])
		} else {
			maxV = nums[i]
		}
	}

	return res
}
