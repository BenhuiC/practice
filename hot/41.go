package hot

func firstMissingPositive(nums []int) int {
	n := len(nums)
	adjust := func(i int) {
		for nums[i] > 0 && nums[i] < n {
			j := nums[i] - 1
			if i == j || nums[i] == nums[j] {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		adjust(i)
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
