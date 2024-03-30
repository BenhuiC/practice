package array

func isMonotonic(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	idx := 1
	for idx < n && nums[idx] == nums[idx-1] {
		idx++
	}
	if idx == n {
		return true
	}

	asc := nums[idx] >= nums[idx-1]
	for idx < n {
		tmp := nums[idx] >= nums[idx-1]
		if nums[idx] == nums[idx-1] || tmp == asc {
			idx++
			continue
		}
		return false
	}
	return true
}
