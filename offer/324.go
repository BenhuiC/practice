package offer

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 1; i < n-1; i += 2 {
		j := i
		for j < n && nums[j] == nums[i] {
			j++
		}
		if j < n {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
