package hot

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := range dp {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
