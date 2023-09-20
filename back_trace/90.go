package back_trace

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	var backtrack func(ary []int, idx int)
	backtrack = func(ary []int, idx int) {
		res = append(res, append([]int{}, ary...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			ary = append(ary, nums[i])
			backtrack(ary, i+1)
			ary = ary[:len(ary)-1]
		}
	}
	sort.Ints(nums)
	backtrack([]int{}, 0)

	return res
}
