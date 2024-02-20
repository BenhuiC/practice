package hot

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return nil
	}
	sort.Ints(candidates)
	if candidates[0] > target {
		return nil
	}

	res := make([][]int, 0)

	var bk func(ary []int, idx, sum int)
	bk = func(ary []int, idx, sum int) {
		if sum > target {
			return
		} else if sum == target {
			res = append(res, append([]int{}, ary...))
			return
		}
		for i := idx; i < n; i++ {
			t := candidates[i] + sum
			if t > target {
				break
			}
			ary = append(ary, candidates[i])
			bk(ary, i, t)
			ary = ary[:len(ary)-1]
		}
	}

	bk([]int{}, 0, 0)
	return res
}
