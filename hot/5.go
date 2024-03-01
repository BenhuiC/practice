package hot

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	maxLen := 0
	startIdx := 0

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				startIdx = i
			}
		}
	}
	return s[startIdx : startIdx+maxLen]
}
