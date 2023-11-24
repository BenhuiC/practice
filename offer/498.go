package offer

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	if m == 0 {
		return nil
	}
	n := len(mat[0])
	res := make([]int, 0, m*n)
	i, j := 0, 0
	order := false
	for i < m && j < n {
		res = append(res, mat[i][j])
		if !order {
			if i == 0 {
				if j < n-1 {
					j++
				} else {
					i++
				}
				order = !order
			} else if j == n-1 {
				i++
				order = !order
			} else {
				i--
				j++
			}
		} else {
			if j == 0 {
				if i < m-1 {
					i++
				} else {
					j++
				}
				order = !order
			} else if i == m-1 {
				j++
				order = !order
			} else {
				i++
				j--
			}

		}
	}

	return res
}
