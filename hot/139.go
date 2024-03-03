package hot

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n < 1 {
		return true
	}
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, w := range wordDict {
			lw := len(w)
			if lw > i {
				continue
			}
			if dp[i-lw] && w == s[i-lw:i] {
				dp[i] = true
				continue
			}
		}
	}

	return dp[n]
}
