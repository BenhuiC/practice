package offer

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		if nums[i-1]-nums[i] != 1 {
			return false
		}
		if i+1 < n && nums[i+1] < nums[i] {
			return false
		}
	}

	return true
}
