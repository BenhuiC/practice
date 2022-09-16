package partice

func maxProduct152(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	imin, imax, res152 := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mi := imin
		ma := imax
		imin = min(ma*nums[i], min(mi*nums[i], nums[i]))
		imax = max(ma*nums[i], max(mi*nums[i], nums[i]))
		res152 = max(res152, imax)
	}

	return res152
}
