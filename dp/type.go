package dp

func Mins(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
