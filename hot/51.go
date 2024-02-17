package hot

import "strings"

func solveNQueens(n int) [][]string {
	genGrid := func(idxs []int) []string {
		g := make([]string, 0, n)
		for _, v := range idxs {
			s := strings.Repeat(".", v) + "Q" + strings.Repeat(".", n-v-1)
			g = append(g, s)
		}
		return g
	}

	res := make([][]string, 0)
	col := make([]bool, n)
	x := make(map[int]bool)
	y := make(map[int]bool)
	var bk func(idxs []int)
	bk = func(idxs []int) {
		if len(idxs) == n {
			res = append(res, genGrid(idxs))
			return
		}
		r := len(idxs)
		for i := 0; i < n; i++ {
			if col[i] || x[r-i] || y[r+i] {
				continue
			}
			x[r-i] = true
			y[r+i] = true
			col[i] = true
			idxs = append(idxs, i)

			bk(idxs)

			x[r-i] = false
			col[i] = false
			y[r+i] = false
			idxs = idxs[:len(idxs)-1]
		}
	}

	bk([]int{})
	return res
}
