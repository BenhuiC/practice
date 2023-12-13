package offer

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var mod int = 10e9 + 7
	dp := make([][][]int, maxMove+1)
	for i := 0; i < maxMove+1; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	for i := 1; i <= maxMove; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				up, down, left, right := 1, 1, 1, 1
				if j != 0 {
					up = dp[i-1][j-1][k]
				}
				if j != m-1 {
					down = dp[i-1][j+1][k]
				}
				if k != 0 {
					left = dp[i-1][j][k-1]
				}
				if k != n-1 {
					right = dp[i-1][j][k+1]
				}
				dp[i][j][k] += up + down + left + right
				dp[i][j][k] %= mod
			}
		}
	}

	return dp[maxMove][startRow][startColumn]
}
