package partice

func containsNearbyDuplicate(nums []int, k int) bool {
	res219 := false
	m := make(map[int]int)
	for i := 0; i <= k && i < len(nums); i++ {
		m[nums[i]]++
	}
	if len(m) < k+1 {
		return true
	}
	left, right := 0, k+1
	for right < len(nums) {
		m[nums[right]]++
		m[nums[left]]--
		for _, v := range m {
			if v >= 2 {
				return true
			}
		}
		left++
		right++
	}

	return res219
}
