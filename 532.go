package partice

import "sort"

func findPairs(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	l, r := 0, 1
	for r < len(nums) {
		if r <= l || nums[r]-nums[l] < k {
			r++
			continue
		}
		if nums[r]-nums[l] > k {
			l++
			continue
		}
		if nums[r]-nums[l] == k {
			res++
			l++
			for l < len(nums) && nums[l] == nums[l-1] {
				l++
			}
		}
	}

	return res
}
