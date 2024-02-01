package hot

func moveZeroes(nums []int) {
	n := len(nums)
	zeroIdx, oneIdx := 0, 0
	for zeroIdx < n && oneIdx < n {
		for zeroIdx < n && nums[zeroIdx] != 0 {
			zeroIdx++
		}
		oneIdx = zeroIdx + 1
		for oneIdx < n && nums[oneIdx] == 0 {
			oneIdx++
		}
		if oneIdx < n {
			nums[zeroIdx], nums[oneIdx] = nums[oneIdx], nums[zeroIdx]
			zeroIdx++
		}
	}
}
