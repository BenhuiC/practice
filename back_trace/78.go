package back_trace

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	var backtrack func(ary []int, idx int)
	backtrack = func(ary []int, idx int) {
		res = append(res, append([]int{}, ary...))
		for i := idx; i < n; i++ {
			ary = append(ary, nums[i])
			backtrack(ary, i+1)
			ary = ary[:len(ary)-1]
		}
	}
	backtrack([]int{}, 0)

	return res
}
