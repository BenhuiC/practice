package offer

func findLengthOfLCIS(nums []int) int {
	l, r := 0, 0
	n := len(nums)
	res := 0
	for l < n {
		for r = l + 1; r < n && nums[r] > nums[l]; r++ {
		}
		if res < r-l {
			res = r - l
		}
		l = r
	}
	return res
}
