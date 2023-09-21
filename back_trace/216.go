package back_trace

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtrace func(ary []int, nu, sum int)
	backtrace = func(ary []int, nu, sum int) {
		if len(ary) == k && sum == n {
			res = append(res, append([]int{}, ary...))
			return
		}
		if len(ary) >= k || sum > n {
			return
		}
		for i := nu; i <= 9 && i < (n-sum)/(k-len(ary))+1; i++ {
			ary = append(ary, i)
			backtrace(ary, i+1, sum+i)
			ary = ary[:len(ary)-1]
		}
	}
	backtrace([]int{}, 1, 0)

	return res
}
