package array

func maxWidthRamp(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	res := 0
	for l := n - 2; l >= 0; l-- {
		r := n - 1
		for r > l && nums[r] < nums[l] {
			r--
		}
		if r-l > res {
			res = r - l
		}
	}

	return res
}
