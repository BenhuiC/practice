package back_trace

func countArrangement(n int) int {
	res := 0
	used := make([]bool, n)
	var bk func(idx int)
	bk = func(idx int) {
		if idx > n {
			res++
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i-1] && (i%idx == 0 || idx%i == 0) {
				used[i-1] = true
				bk(idx + 1)
				used[i-1] = false
			}
		}
	}
	bk(1)
	return res
}
