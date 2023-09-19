package back_trace

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	var backtrack func(ary []int, idx int)
	vis := make([]bool, len(nums))
	backtrack = func(ary []int, idx int) {
		if idx == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, ary)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if vis[i] || i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
				continue
			}
			ary = append(ary, nums[i])
			vis[i] = true
			backtrack(ary, idx+1)
			vis[i] = false
			ary = ary[:len(ary)-1]
		}
	}
	sort.Ints(nums)
	backtrack([]int{}, 0)

	return res
}
