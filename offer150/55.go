package offer150

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	step := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= step {
			step = 1
		} else {
			step++
		}
	}
	return step == 1
}
