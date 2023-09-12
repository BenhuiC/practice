package dp

/*
dp[i][j]表示为s[i,j]区间形成回文需要插入多少个字符
*/
func minInsertions(s string) int {
	slen := len(s)
	if slen <= 1 {
		return 0
	}
	dp := make([][]int, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]int, slen)
	}
	for i := slen - 1; i >= 0; i-- {
		for j := i + 1; j < slen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = Min(dp[i+1][j-1]+2, dp[i][j-1]+1, dp[i+1][j]+1)
			}
		}
	}

	return dp[0][slen-1]
}
