package dp

func orderOfLargestPlusSign(n int, mines [][]int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = n
		}
	}
	mp := make(map[int]bool)
	for _, m := range mines {
		mp[m[0]*n+m[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if !mp[i*n+j] {
				count++
			} else {
				count = 0
			}
			dp[i][j] = min(count, dp[i][j])
		}
		count = 0
		for j := n - 1; j >= 0; j-- {
			if !mp[i*n+j] {
				count++
			} else {
				count = 0
			}
			dp[i][j] = min(count, dp[i][j])
		}
	}

	res := 0
	for j := 0; j < n; j++ {
		count := 0
		for i := 0; i < n; i++ {
			if !mp[i*n+j] {
				count++
			} else {
				count = 0
			}
			dp[i][j] = min(count, dp[i][j])
		}
		count = 0
		for i := n - 1; i >= 0; i-- {
			if !mp[i*n+j] {
				count++
			} else {
				count = 0
			}
			dp[i][j] = min(count, dp[i][j])
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
