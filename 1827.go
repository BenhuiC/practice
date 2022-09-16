package partice

func minOperations1827(nums []int) int {
	res := 0
	if len(nums) < 2 {
		return res
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		curVal := nums[i-1] + 1
		res += curVal - nums[i]
		nums[i] = curVal
	}

	return res
}
