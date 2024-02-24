package hot

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	steal, noSteal := nums[0], 0
	for i := 1; i < n; i++ {
		steal, noSteal = noSteal+nums[i], max(noSteal, steal)
	}

	return max(noSteal, steal)
}
