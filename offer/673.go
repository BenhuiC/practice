package offer

func findNumberOfLIS(nums []int) int {
	maxLen := 0
	res := 0
	n := len(nums)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{1, 1}
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				if dp[j][0]+1 > dp[i][0] {
					dp[i] = [2]int{dp[j][0] + 1, dp[j][1]}
				} else if dp[j][0]+1 == dp[i][0] {
					dp[i][1] += dp[j][1]
				}
			}
		}
		if dp[i][0] > maxLen {
			maxLen = dp[i][0]
			res = dp[i][1]
		} else if dp[i][0] == maxLen {
			res += dp[i][1]
		}
	}

	return res
}
