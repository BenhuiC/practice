package back_trace

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	res := 0
	col := make([]bool, n)
	x := make(map[int]bool)
	y := make(map[int]bool)
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			if col[i] || x[r-i] || y[r+i] {
				continue
			}

			x[r-i] = true
			y[r+i] = true
			col[i] = true

			backtrack(r + 1)

			x[r-i] = false
			col[i] = false
			y[r+i] = false
		}
	}
	backtrack(0)

	return res
}
