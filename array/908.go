package array

import "sort"

func smallestRangeI(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)

	mi, mx := nums[0], nums[n-1]
	mi += k
	mx -= k

	if mi >= mx {
		return 0
	}
	return mx - mi
}
