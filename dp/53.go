package dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + Max(0, nums[i-1])
		res = Max(res, nums[i])
	}

	return res
}
