package partice

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	m, n := len(text1)+1, len(text2)+1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			if text1[x-1] == text2[y-1] {
				dp[x][y] = dp[x-1][y-1] + 1
			} else {
				dp[x][y] = Max(dp[x-1][y], dp[x][y-1])
			}
		}
	}
	return dp[m-1][n-1]
}
