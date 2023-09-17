package greedy

func jump(nums []int) int {
	n := len(nums)
	steps := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < len(nums) && j <= nums[i]+i; j++ {
			if steps[j] == 0 {
				steps[j] = steps[i] + 1
			} else {
				steps[j] = Min(steps[j], steps[i]+1)
			}
		}
	}

	return steps[n-1]
}
