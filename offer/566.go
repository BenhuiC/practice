package offer

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	if m == 0 {
		return nil
	}
	n := len(mat[0])
	if m == r || m*n != r*c {
		return mat
	}

	res := make([][]int, r)
	res[0] = make([]int, c)
	rt, ct := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[rt][ct] = mat[i][j]
			ct++
			if ct >= c {
				rt++
				ct = 0
				if rt < r {
					res[rt] = make([]int, c)
				} else {
					return res
				}
			}
		}
	}
	return res
}
