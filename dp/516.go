package dp

/*
dp[x][y]表示为s[x:y]区间的最长回问子串长度
如果s[x]==s[y] 则s[x:y]=s[x+1,y-1]+2
否则s[x:y]=max(s[x,y-1],s[x+1,y])
*/
func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	slen := len(s)
	dp := make([][]int, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]int, slen)
		dp[i][i] = 1
	}

	for i := slen - 1; i >= 0; i-- {
		for j := i + 1; j < slen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][slen-1]
}
