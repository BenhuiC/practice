package back_trace

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var backtrack func(ary []int, idx, sum int)
	backtrack = func(ary []int, idx, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, ary...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			ary = append(ary, candidates[i])
			backtrack(ary, i+1, sum+candidates[i])
			ary = ary[:len(ary)-1]
		}
	}
	sort.Ints(candidates)
	backtrack([]int{}, 0, 0)

	return res
}
