package back_trace

import "strings"

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	res := make([][]string, 0)
	col := make([]bool, n)
	x := make(map[int]bool)
	y := make(map[int]bool)
	var backtrack func(ary []string)
	backtrack = func(ary []string) {
		if len(ary) == n {
			res = append(res, append([]string{}, ary...))
			return
		}
		r := len(ary)
		for i := 0; i < n; i++ {
			if col[i] || x[r-i] || y[r+i] {
				continue
			}

			cur := strings.Repeat(".", i) + "Q" + strings.Repeat(".", n-i-1)
			ary = append(ary, cur)
			x[r-i] = true
			y[r+i] = true
			col[i] = true

			backtrack(ary)

			x[r-i] = false
			col[i] = false
			y[r+i] = false
			ary = ary[:len(ary)-1]
		}
	}
	backtrack([]string{})

	return res
}
