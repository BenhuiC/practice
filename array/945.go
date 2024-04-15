package array

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		to := nums[i-1] + 1
		res += to - nums[i]
		nums[i] = to
	}

	return res
}
