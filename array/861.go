package array

func matrixScore(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	res := 1 << (n - 1) * m
	for j := 1; j < n; j++ {
		oneCnt := 0
		for _, row := range grid {
			// 如果此列的值等于列首的值，则该值比为1
			if row[j] == row[0] {
				oneCnt++
			}
		}
		// 翻转此列
		if oneCnt < m-oneCnt {
			oneCnt = m - oneCnt
		}
		res += 1 << (n - j - 1) * oneCnt
	}

	return res
}
