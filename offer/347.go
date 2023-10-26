package offer

import "sort"

func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)
	ary := make([][2]int, 0, len(nums))
	ary = append(ary, [2]int{nums[0], 1})
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ary[len(ary)-1][1]++
		} else {
			ary = append(ary, [2]int{nums[i], 1})
		}
	}
	sort.Slice(ary, func(i, j int) bool {
		return ary[i][1] > ary[j][1]
	})
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, ary[i][0])
	}
	return res
}
