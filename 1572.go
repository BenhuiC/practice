package partice

func diagonalSum(mat [][]int) int {
	n := len(mat)
	res1572 := 0
	x, y := 0, n-1
	for i := 0; i < n; i++ {
		if x == y {
			res1572 += mat[x][i]
		} else {
			res1572 += mat[x][i]
			res1572 += mat[y][i]
		}
		x++
		y--
	}

	return res1572
}
