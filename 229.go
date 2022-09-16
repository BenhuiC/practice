package partice

import "sort"

func majorityElement(nums []int) []int {
	res229 := make([]int, 0)
	sort.Ints(nums)
	nlen := len(nums)
	limit := nlen / 3
	i, j := 0, 0
	for i < nlen && j < nlen {
		for j < nlen && nums[j] == nums[i] {
			j++
		}
		if j-i > limit {
			res229 = append(res229, nums[i])
		}
		i = j
	}
	return res229
}
