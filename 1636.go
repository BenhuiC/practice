package partice

import "sort"

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	sort.Slice(&nums, func(i, j int) bool {
		if m[nums[i]] < m[nums[j]] {
			return true
		} else if m[nums[i]] == m[nums[j]] {
			return nums[i] > nums[j]
		}

		return false
	})
	return nums
}
