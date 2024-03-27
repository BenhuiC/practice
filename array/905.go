package array

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		for l < n && nums[l]&1 == 0 {
			l++
		}
		for r >= 0 && nums[r]&1 == 1 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}
