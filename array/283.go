package array

func moveZeroes(nums []int) {
	zeroIdx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zeroIdx == -1 {
			zeroIdx = i
		} else if nums[i] != 0 && zeroIdx != -1 {
			nums[i], nums[zeroIdx] = nums[zeroIdx], nums[i]
			zeroIdx++
		}
	}
}
