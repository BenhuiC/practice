package hot

import "sort"

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	res := 1
	sort.Ints(nums)
	l := 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			l++
		} else {
			l = 1
		}
		if l > res {
			res = l
		}
	}

	return res
}
