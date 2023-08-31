package interview

func setZeroes(matrix [][]int) {
	c := make(map[int]struct{})
	r := make(map[int]struct{})
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				c[i] = struct{}{}
				r[j] = struct{}{}
			}
		}
	}

	for k := range c {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}
	for k := range r {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}
