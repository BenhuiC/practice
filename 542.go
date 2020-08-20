package partice

func updateMatrix(matrix [][]int) [][]int {
	result := make([][]int, len(matrix))
	for i := 0; i < len(matrix[0]); i++ {
		result[i] = make([]int, len(matrix[i]))
	}

	return result
}
