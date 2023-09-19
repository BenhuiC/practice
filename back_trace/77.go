package back_trace

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrack func(nu int, ary []int)
	backtrack = func(nu int, ary []int) {
		if len(ary) == k {
			res = append(res, append([]int{}, ary...))
			return
		}
		// 剪枝
		for i := nu; i <= n-(k-len(ary))+1; i++ {
			ary = append(ary, i)
			backtrack(i+1, ary)
			ary = ary[:len(ary)-1]
		}
	}
	backtrack(1, []int{})

	return res
}
