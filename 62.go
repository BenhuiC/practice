package partice

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	mp := make([][]int, m)
	for i := 0; i < m; i++ {
		mp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		mp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		mp[0][j] = 1
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			mp[x][y] = mp[x-1][y] + mp[x][y-1]
		}
	}
	return mp[m-1][n-1]
}
