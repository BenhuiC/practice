package partice

func maxProduct152(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	imin, imax, res152 := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mi := imin
		ma := imax
		imin = Min(ma*nums[i], Min(mi*nums[i], nums[i]))
		imax = Max(ma*nums[i], Max(mi*nums[i], nums[i]))
		res152 = Max(res152, imax)
	}

	return res152
}
