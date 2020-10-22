package partice

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i], matrix[j][n-i-1] = matrix[j][n-i-1], matrix[j][i]
		}
	}
}
