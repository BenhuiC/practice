package offer150

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	step, next, maxDis := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxDis = max(maxDis, i+nums[i])
		if i == next {
			step++
			next = maxDis
		}
	}

	return step
}
