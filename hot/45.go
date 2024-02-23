package hot

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	var start, end, max, step int
	for end < n-1 {
		for i := start; i <= end; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		}
		start = end + 1
		end = max
		step++
	}

	return step
}
