package partice

func gridGame(grid [][]int) int64 {
	var sum1, sum2 int64
	n := len(grid[0])
	if n == 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		sum1 += int64(grid[0][i])
		sum2 += int64(grid[1][i])
	}
	res2017 := sum1 + sum2
	var tmpSum1, tmpSum2 int64
	for j := 0; j < n; j++ {
		tmpSum1 += int64(grid[0][j])
		if tmpSum1+sum2-tmpSum2 > maxInt64(tmpSum2, sum1-tmpSum1) {
			res2017 = minInt64(res2017, maxInt64(tmpSum2, sum1-tmpSum1))
		}
		tmpSum2 += int64(grid[1][j])
	}
	return res2017
}

func maxInt64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func minInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
