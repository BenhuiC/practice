package offer

func findMaxForm(strs []string, m int, n int) int {
	nl := len(strs)
	cnt := make([][2]int, nl)
	for i, str := range strs {
		for _, c := range str {
			cnt[i][c-'0']++
		}
	}

	dp := make([][][]int, nl+1)
	for i := 0; i <= nl; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i, c := range cnt {
		zero, one := c[0], c[1]
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zero && k >= one {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zero][k-one]+1)
				}
			}
		}
	}

	return dp[nl][m][n]
}
