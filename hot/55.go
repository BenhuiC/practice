package hot

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	step := 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= step {
			step = 1
		} else {
			step++
		}
	}

	return step == 1
}
