package partice

// 连续子数组
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	var res413 int
	d, t := nums[1]-nums[0], 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == d {
			t++
		} else {
			d = nums[i] - nums[i-1]
			t = 0
		}
		res413 += t
	}

	return res413
}
