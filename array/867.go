package array

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		ary := make([]int, 0, m)
		for j := 0; j < m; j++ {
			ary = append(ary, matrix[j][i])
		}

		res[i] = ary
	}

	return res
}
