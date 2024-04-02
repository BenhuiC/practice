package array

import "sort"

func smallestRangeII(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	sort.Ints(nums)

	res := nums[n-1] - nums[0]
	for i := 0; i < n-1; i++ {
		a, b := nums[i], nums[i+1]
		high := max(nums[n-1]-k, a+k)
		low := min(nums[0]+k, b-k)
		res = min(res, high-low)
	}

	return res
}
